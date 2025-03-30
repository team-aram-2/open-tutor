package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"open-tutor/internal/services/db"
	"open-tutor/middleware"
	"open-tutor/stripe_client"
	"open-tutor/util"

	"github.com/lib/pq"
	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/stripe/stripe-go/v81"
)

type TutorResponse struct {
	Info         Tutor        `json:"info"`
	RatingScores RatingScores `json:"ratingScores"`
	RatingCount  int          `json:"ratingCount"`
}

func (t *OpenTutor) SignUpAsTutor(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("error parsing form data: %v", err))
		return
	}

	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo.UserID == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Check if user exists
	var email, userFirstName, userLastName string
	var currentRoleMask util.RoleMask
	err = db.GetDB().QueryRow("SELECT email, first_name, last_name, role_mask FROM users WHERE user_id = $1", authInfo.UserID).Scan(&email, &userFirstName, &userLastName, &currentRoleMask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Database error: %s\n", err)
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User with ID:{} does not exist\n")
		return
	}

	// Check if user exists as tutor already
	var tutorExists bool
	err = db.GetDB().QueryRow("SELECT EXISTS(SELECT 1 FROM tutors WHERE user_id = $1)", authInfo.UserID).Scan(&tutorExists)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Database error: %s\n", err)
		return
	}
	if tutorExists {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "User with ID:{} is already registered as a tutor.\n")
		return
	}

	// Update user role to include Tutor role
	err = db.GetDB().QueryRow("SELECT role_mask FROM users WHERE user_id = $1", authInfo.UserID).Scan(&currentRoleMask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Database error: %s\n", err)
		return
	}

	// Update the user's role in the database
	_, err = db.GetDB().Exec("UPDATE users SET role_mask = $1 WHERE user_id = $2", currentRoleMask, authInfo.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to update user's role: %v\n", err)
		return
	}

	// Update the session token to reflect the new role
	err = applySessionTokenForUserId(authInfo.UserID, false, currentRoleMask, w)
	if err != nil {
		fmt.Printf("failed to apply session token: %v\n", err)
		sendError(w, http.StatusInternalServerError, "failed to apply session token")
		return
	}

	// Add the Tutor role to the current role mask
	currentRoleMask.Add(util.Tutor)

	_, insertErr := db.GetDB().Exec("INSERT INTO tutors (user_id) VALUES ($1)", authInfo.UserID)
	if insertErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", insertErr)
		return
	}

	requestIp := strings.Split(r.RemoteAddr, ":")[0]

	// Create Stripe account for tutor //
	params := &stripe.AccountParams{
		Type:    stripe.String("custom"),
		Country: stripe.String("US"),
		Email:   stripe.String(email),
		Capabilities: &stripe.AccountCapabilitiesParams{
			CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
				Requested: stripe.Bool(true),
			},
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
		ExternalAccount: &stripe.AccountExternalAccountParams{
			Country:       stripe.String("US"),
			Currency:      stripe.String("usd"),
			AccountNumber: stripe.String(r.FormValue("bank_account_number")),
			RoutingNumber: stripe.String(r.FormValue("bank_routing_number")),
		},
		TOSAcceptance: &stripe.AccountTOSAcceptanceParams{
			Date: stripe.Int64(time.Now().Unix()), // Unix timestamp
			IP:   stripe.String(requestIp),
		},
		Settings: &stripe.AccountSettingsParams{
			Payments: &stripe.AccountSettingsPaymentsParams{
				StatementDescriptor: stripe.String("OPENTUTOR"),
			},
		},
		BusinessType: stripe.String("individual"),
		BusinessProfile: &stripe.AccountBusinessProfileParams{
			Name: stripe.String("OpenTutor"),
			URL:  stripe.String("https://opentutor.com"),
			MCC:  stripe.String("8299"), // mcc for education services
		},
		Individual: &stripe.PersonParams{
			FirstName: stripe.String(userFirstName),
			LastName:  stripe.String(userLastName),
			Email:     stripe.String(email),
			Phone:     stripe.String(r.FormValue("phone")),
			Address: &stripe.AddressParams{
				Line1:      stripe.String(r.FormValue("address_line1")),
				City:       stripe.String(r.FormValue("address_city")),
				State:      stripe.String(r.FormValue("address_state")),
				PostalCode: stripe.String(r.FormValue("address_postalcode")),
				Country:    stripe.String("US"),
			},
			DOB: &stripe.PersonDOBParams{
				Day:   stripe.Int64(2),
				Month: stripe.Int64(5),
				Year:  stripe.Int64(2002),
			},
			IDNumber: stripe.String(r.FormValue("ssn")),
		},
	}
	result, err := stripe_client.GetClient().Accounts.New(params)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("failed to create Stripe account for tutor: %v", err))
		return
	}

	_, err = db.GetDB().Exec("UPDATE tutors SET stripe_account_id = $1 WHERE user_id = $2", result.ID, authInfo.UserID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("failed to save Stripe account id to tutor: %v", err))
		return
	}
	w.Header().Set("Location", r.Header.Get("Origin")+"/")
	w.WriteHeader(http.StatusPermanentRedirect)
}

func (t *OpenTutor) GetTutors(w http.ResponseWriter, r *http.Request, params GetTutorsParams) {
	// Base query for tutor details
	query := `
	SELECT u.user_id, u.first_name, u.last_name, u.email, u.signed_up_at,
		   t.total_hours, ARRAY_REMOVE(ARRAY_AGG(ts.skill_id), NULL) AS skills,
		   COALESCE(AVG(r.communication), 0), COALESCE(AVG(r.knowledge), 0),
		   COALESCE(AVG(r.overall), 0), COALESCE(AVG(r.professionalism), 0),
		   COALESCE(AVG(r.punctuality), 0), COUNT(r.user_id)
	FROM tutors t
	INNER JOIN users u ON t.user_id = u.user_id
	LEFT JOIN tutor_skills ts ON t.user_id = ts.tutor_id
	LEFT JOIN ratings r ON t.user_id = r.user_id AND r.rating_type = 'tutor'
`
	var args []interface{}
	conditions := []string{}
	havingConditions := []string{}

	argIndex := 1 // PostgreSQL parameters start from $1

	// Handle Skills Include Filter
	if params.SkillsInclude != nil && len(*params.SkillsInclude) > 0 {
		conditions = append(conditions, fmt.Sprintf("t.user_id IN (SELECT tutor_id FROM tutor_skills WHERE skill_id = ANY($%d))", argIndex))
		args = append(args, pq.Array(*params.SkillsInclude))
		argIndex++
	}

	// Apply conditions before GROUP BY
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// GROUP BY (needed for aggregation)
	query += `
	GROUP BY u.user_id, u.first_name, u.last_name, u.email, u.signed_up_at, t.total_hours, u.account_locked
`

	// Move minRating filter to HAVING clause
	if params.MinRating != nil {
		havingConditions = append(havingConditions, fmt.Sprintf("COALESCE(AVG(r.overall), 0) >= $%d", argIndex))
		args = append(args, *params.MinRating)
		argIndex++
	}

	// Apply HAVING clause if needed
	if len(havingConditions) > 0 {
		query += " HAVING " + strings.Join(havingConditions, " AND ")
	}

	// ORDER, LIMIT, and OFFSET
	query += fmt.Sprintf(" ORDER BY AVG(r.overall) DESC LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, params.PageSize, params.PageSize*params.PageIndex)

	// Execute query
	rows, err := db.GetDB().Query(query, args...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Parse results
	var tutorResponses []TutorResponse
	for rows.Next() {
		var tutor Tutor
		var skills []string
		var ratingScores RatingScores
		var ratingCount int

		err := rows.Scan(
			&tutor.UserId, &tutor.FirstName, &tutor.LastName, &tutor.Email, &tutor.SignedUpAt,
			&tutor.TotalHours, pq.Array(&skills),
			&ratingScores.Communication, &ratingScores.Knowledge, &ratingScores.Overall,
			&ratingScores.Professionalism, &ratingScores.Punctuality, &ratingCount,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse tutors: %v", err), http.StatusInternalServerError)
			return
		}

		tutor.Skills = &skills

		tutorResponse := TutorResponse{
			Info:         tutor,
			RatingScores: ratingScores,
			RatingCount:  ratingCount,
		}

		tutorResponses = append(tutorResponses, tutorResponse)
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tutorResponses)
	if err != nil {
		log.Println("Error: ", err)
	}
}

func (t *OpenTutor) GetTutorById(w http.ResponseWriter, r *http.Request, tutorId openapi_types.UUID) {
	tutor := &Tutor{
		Skills: new([]string),
	}
	selectErr := db.GetDB().QueryRow(
		`
		SELECT
			first_name,
			last_name,
			signed_up_at,
			total_hours,
			tutors.user_id
		FROM users
		INNER JOIN tutors ON users.user_id = tutors.user_id
		WHERE tutors.user_id = $1
		`,
		tutorId,
	).Scan(
		&tutor.FirstName,
		&tutor.LastName,
		&tutor.SignedUpAt,
		&tutor.TotalHours,
		&tutor.UserId,
	)

	if selectErr != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s\n", selectErr)
		return
	}
	// Get skills arr from tutor_skills
	var skills sql.NullString
	selectErr = db.GetDB().QueryRow(`
		SELECT ARRAY_AGG(skill_id) AS skills
		FROM tutor_skills
		WHERE tutor_skills.tutor_id = $1
	`, tutorId,
	).Scan(&skills)

	if selectErr != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s\n", selectErr)
		return
	}

	var skillList []string
	if skills.Valid { // Only parse if NOT NULL
		err := json.Unmarshal([]byte(skills.String), &skillList)
		if err != nil {
			log.Printf("Failed to unmarshal skills: %v", err)
			skillList = []string{} // Set an empty slice if NULL
		}
	} else {
		skillList = []string{} // Set an empty slice if NULL
	}

	*tutor.Skills = skillList

	// Get all ratings for a tutor
	rows, selectErr := db.GetDB().Query(`
		SELECT
		 	ratings.overall,
			ratings.professionalism,
			ratings.knowledge,
			ratings.communication,
			ratings.punctuality
		FROM ratings
		WHERE ratings.user_id = $1 AND ratings.rating_type = 'tutor'
	`, tutorId,
	)
	if selectErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", selectErr)
		return
	}

	// aggregate the stats
	var aggregate RatingScores
	var ratingCount int = 0
	for rows.Next() {
		var scores RatingScores
		err := rows.Scan(&scores.Overall, &scores.Professionalism, &scores.Knowledge, &scores.Communication, &scores.Punctuality)
		if err != nil {
			log.Println("Error scanning rating:", err)
			continue
		}
		aggregate.Overall += scores.Overall
		aggregate.Professionalism += scores.Professionalism
		aggregate.Knowledge += scores.Knowledge
		aggregate.Communication += scores.Communication
		aggregate.Punctuality += scores.Punctuality
		ratingCount += 1
	}
	defer rows.Close()

	if ratingCount > 0 {
		aggregate.Overall = aggregate.Overall / ratingCount
		aggregate.Professionalism = aggregate.Professionalism / ratingCount
		aggregate.Knowledge = aggregate.Knowledge / ratingCount
		aggregate.Communication = aggregate.Communication / ratingCount
		aggregate.Punctuality = aggregate.Punctuality / ratingCount
	}

	response := TutorResponse{
		Info:         *tutor,
		RatingScores: aggregate,
		RatingCount:  ratingCount,
	}
	fmt.Println(response)

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println("Error encoding json: ", err)
	}
}

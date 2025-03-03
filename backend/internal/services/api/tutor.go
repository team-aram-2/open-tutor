package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"open-tutor/internal/services/db"
	"open-tutor/middleware"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type TutorResponse struct {
	Info         Tutor        `json:"info"`
	RatingScores RatingScores `json:"ratingScores"`
	RatingCount  int          `json:"ratingCount"`
}

func (t *OpenTutor) SignUpAsTutor(w http.ResponseWriter, r *http.Request) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	_, insertErr := db.GetDB().Exec("INSERT INTO tutors (user_id) VALUES ($1)", authInfo.UserID)
	if insertErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", insertErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (t *OpenTutor) GetTutorById(w http.ResponseWriter, r *http.Request, tutorId openapi_types.UUID) {
	tutor := &Tutor{}
	selectErr := db.GetDB().QueryRow(
		`
		SELECT
			first_name,
			last_name,
			signed_up_at,
			total_hours
		FROM users
		INNER JOIN tutors ON users.user_id = tutors.user_id
		WHERE user_id = $1
		`,
		tutorId,
	).Scan(
		&tutor.FirstName,
		&tutor.LastName,
		&tutor.SignedUpAt,
		&tutor.TotalHours,
	)

	if selectErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", selectErr)
		return
	}
	// Get skills arr from tutor_skills
	var skills []string
	selectErr = db.GetDB().QueryRow(`
		SELECT ARRAY_AGG(skill_id) AS skills
		FROM tutor_skills
		WHERE tutor_skills.tutor_id = $1
	`, tutorId,
	).Scan(&skills)

	if selectErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", selectErr)
		return
	}

	*tutor.Skills = skills

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

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

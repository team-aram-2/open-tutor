package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"open-tutor/internal/services/db"
	"open-tutor/middleware"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

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
	selectErr := db.GetDB().QueryRow(`
		SELECT [
			first_name,
			last_name,
			signed_up_at,
		]
		FROM users
		INNER JOIN tutors ON users.user_id = tutors.user_id
	`, tutorId).Scan(
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
	// TODO: 1. Need to get SKILLS from tutor_skills as UUID string array

	// TODO: 2. Get all ratings for a tutor and aggregate the stats to return as part of the struct

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tutor)
}

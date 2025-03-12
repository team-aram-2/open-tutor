package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"open-tutor/internal/services/db"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func checkUser(userId openapi_types.UUID) (bool, error) {
	var exist bool
	err := db.GetDB().QueryRow(`
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE user_id = $1
		)
	`, userId).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (t *OpenTutor) DeleteUserById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID) {
	var exist bool
	var err error
	exist, err = checkUser(userId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if !exist {
		sendError(w, http.StatusNotFound, "User not found")
		return
	}

	_, deleteErr := db.GetDB().Exec(`
		DELETE FROM users
		WHERE user_id = $1
		`, userId)

	if deleteErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", deleteErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (t *OpenTutor) GetUserById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID) {
	var exist bool
	var err error
	exist, err = checkUser(userId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if !exist {
		sendError(w, http.StatusNotFound, "User not found")
		return
	}

	user := &User{}
	selectErr := db.GetDB().QueryRow(`
		SELECT *
		FROM users
		WHERE user_id = $1
	`, userId).Scan(
		&user.UserId,
		&user.Email,
		&user.SignedUpAt,
		&user.FirstName,
		&user.LastName,
		&user.AccountLocked,
	)

	if selectErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", selectErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}

func (t *OpenTutor) UpdateUserById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID) {
	var exist bool
	var err error
	exist, err = checkUser(userId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if !exist {
		sendError(w, http.StatusNotFound, "User not found")
		return
	}

	var user User
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for user")
		return
	}

	updateErr := db.GetDB().QueryRow(`
		UPDATE users
		SET first_name = $1, last_name = $2, email = $3
		WHERE user_id = $4
		RETURNING user_id, email, first_name, last_name;
	`, user.FirstName, user.LastName, user.Email, userId).Scan(
		&updatedUser.UserId,
		&updatedUser.Email,
		&updatedUser.FirstName,
		&updatedUser.LastName,
	)
	if updateErr != nil {
		sendError(w, http.StatusInternalServerError, "Database error")
		fmt.Fprintf(w, "%s\n", updateErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(updatedUser)
}

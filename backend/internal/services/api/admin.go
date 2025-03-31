package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"open-tutor/internal/services/db"
	"open-tutor/middleware"
	"open-tutor/util"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (t *OpenTutor) UpdateUserRole(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo == nil || authInfo.UserID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ensure the authenticated user has admin rights
	if !authInfo.RoleMask.Has(util.Admin) {
		http.Error(w, "Forbidden: Admin access required", http.StatusForbidden)
		return
	}

	// Parse the request body for the new role
	var roleUpdate struct {
		Role util.RoleMask `json:"role"`
	}
	err := json.NewDecoder(r.Body).Decode(&roleUpdate)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Ensure the role is valid
	valid := false
	for _, validRole := range util.Roles() {
		if validRole == roleUpdate.Role {
			valid = true
			break
		}
	}
	if !valid {
		http.Error(w, "Invalid role provided", http.StatusBadRequest)
		return
	}

	// Update the user's role in the database
	_, err = db.GetDB().Exec(`
		UPDATE users
		SET role_mask = $1
		WHERE user_id = $2
	`, roleUpdate.Role, userId)
	if err != nil {
		http.Error(w, "Failed to update role", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User role updated successfully")
}

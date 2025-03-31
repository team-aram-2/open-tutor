package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"open-tutor/internal/services/db"
	"open-tutor/middleware"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// GetCategories retrieves all academic categories
func (t *OpenTutor) GetCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := db.GetDB().Query("SELECT id, name FROM academic_categories")
	if err != nil {
		http.Error(w, "Failed to retrieve categories", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []AcademicCategory
	for rows.Next() {
		var category AcademicCategory
		if err := rows.Scan(&category.Id, &category.Name); err != nil {
			http.Error(w, "Error scanning category", http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(categories)
}

// CreateCategory adds a new academic category (Admin only)
func (t *OpenTutor) CreateCategory(w http.ResponseWriter, r *http.Request) {
	authInfo := middleware.GetAuthenticationInfo(r)
	// TODO: NATHAN: I NEED TO FINALIZE IMPLEMENTING RBAC ALREADY FFS
	if authInfo.UserID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var requestBody CreateCategoryJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Insert new category
	newID := uuid.New()
	_, err := db.GetDB().Exec("INSERT INTO academic_categories (id, name) VALUES ($1, $2)", newID, requestBody.Name)
	if err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	category := AcademicCategory{Id: newID, Name: requestBody.Name}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		http.Error(w, "Error scanning category", http.StatusInternalServerError)
		return
	}
}

// GetCategory retrieves details of a specific academic category
func (t *OpenTutor) GetCategory(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var category AcademicCategory
	err := db.GetDB().QueryRow("SELECT id, name FROM academic_categories WHERE id = $1", id).Scan(&category.Id, &category.Name)

	if err == sql.ErrNoRows {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to retrieve category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		http.Error(w, "Error scanning category", http.StatusInternalServerError)
		return
	}
}

// UpdateCategory updates an academic category (Admin only)
func (t *OpenTutor) UpdateCategory(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	authInfo := middleware.GetAuthenticationInfo(r)
	// TODO: NATHAN: I NEED TO FINALIZE IMPLEMENTING RBAC ALREADY FFS
	if authInfo.UserID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var requestBody UpdateCategoryJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update category
	result, err := db.GetDB().Exec("UPDATE academic_categories SET name = $1 WHERE id = $2", requestBody.Name, id)
	if err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	category := AcademicCategory{Id: id, Name: requestBody.Name}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		http.Error(w, "Error scanning category", http.StatusInternalServerError)
		return
	}
}

// DeleteCategory removes an academic category (Admin only)
func (t *OpenTutor) DeleteCategory(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	authInfo := middleware.GetAuthenticationInfo(r)
	// TODO: NATHAN: I NEED TO FINALIZE IMPLEMENTING RBAC ALREADY FFS
	if authInfo.UserID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Delete category
	result, err := db.GetDB().Exec("DELETE FROM academic_categories WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

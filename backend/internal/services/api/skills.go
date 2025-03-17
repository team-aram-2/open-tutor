package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"open-tutor/internal/services/db"
	"open-tutor/middleware"

	"github.com/google/uuid"
	"github.com/lib/pq"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// NEEDS TO BE CONFIGURABLE BY AN ADMINISTRATOR AND SERIALIZED TO A SETTINGS FILE SOMEWHERE
const (
	N_QUIZ_QUESTIONS = 20
	PASSING_RATIO    = 0.8
)

type QuizAttemptResponse struct {
	Success bool `json:"success"`
}

// GetAllSkills retrieves a list of all skills.
// (GET /skills)
func (t *OpenTutor) GetAllSkills(w http.ResponseWriter, r *http.Request) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo.UserID == "" {
		http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// Parse optional "category" query parameter
	categoryID := r.URL.Query().Get("category")
	var rows *sql.Rows
	var err error

	if categoryID != "" {
		rows, err = db.GetDB().Query(
			"SELECT id, name, category FROM skills WHERE category = $1",
			categoryID,
		)
	} else {
		rows, err = db.GetDB().Query("SELECT id, name, category FROM skills")
	}

	if err != nil {
		http.Error(w, `{"error": "Database error"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Process skill records
	var skills []TutorSkill
	for rows.Next() {
		var skill TutorSkill
		if err := rows.Scan(&skill.Id, &skill.Title, &skill.Category); err != nil {
			http.Error(w, `{"error": "Failed to scan skill data"}`, http.StatusInternalServerError)
			return
		}
		skills = append(skills, skill)
	}

	// Check if no skills were found
	if len(skills) == 0 {
		http.Error(w, `{"error": "No skills found matching query"}`, http.StatusNotFound)
		return
	}

	// Return skills as JSON response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(skills)
	if err != nil {
		http.Error(w, "Error scanning category", http.StatusInternalServerError)
		return
	}
}

// GetSkill retrieves information for a specific skill by ID.
// (GET /skills/{id})
func (t *OpenTutor) GetSkill(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var skill TutorSkill
	err := db.GetDB().QueryRow(
		"SELECT id, title, description, category_id FROM skills WHERE id = $1", id,
	).Scan(&skill.Id, &skill.Title, &skill.Description, &skill.Category)

	if err == sql.ErrNoRows {
		http.Error(w, "Skill not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(skill)
	if err != nil {
		http.Error(w, "Error scanning category", http.StatusInternalServerError)
		return
	}
}

// CreateSkill handles the creation of a new skill.
// (POST /skills)
func (t *OpenTutor) CreateSkill(w http.ResponseWriter, r *http.Request) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo.UserID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Decode request body
	var req TutorSkill
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check for duplicate skill title in the same category
	var existingSkillId uuid.UUID
	err := db.GetDB().QueryRow(
		"SELECT id FROM skills WHERE title = $1 AND category_id = $2",
		req.Title, req.Category,
	).Scan(&existingSkillId)

	if err == nil {
		http.Error(w, "A skill with this title already exists in the selected category", http.StatusConflict)
		return
	} else if err != sql.ErrNoRows {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}

	// Insert new skill and return the generated ID
	var skillId uuid.UUID
	err = db.GetDB().QueryRow(
		"INSERT INTO skills (title, description, category_id) VALUES ($1, $2, $3) RETURNING id",
		req.Title, req.Description, req.Category,
	).Scan(&skillId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create skill: %v", err), http.StatusInternalServerError)
		return
	}

	// Return created skill with ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":          skillId,
		"title":       req.Title,
		"description": req.Description,
		"categoryId":  req.Category,
	})
}

// UpdateSkill (POST /skills/{id})
func (t *OpenTutor) UpdateSkill(w http.ResponseWriter, r *http.Request, skillId openapi_types.UUID) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo.UserID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var skill TutorSkill
	if err := json.NewDecoder(r.Body).Decode(&skill); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to start transaction: %v", err), http.StatusInternalServerError)
		return
	}

	// Update skill details
	_, err = tx.Exec(
		"UPDATE available_skills SET category = $1, title = $2, description = $3 WHERE id = $4",
		skill.Category, skill.Title, skill.Description, skillId,
	)
	if err != nil {
		_ = tx.Rollback()
		http.Error(w, fmt.Sprintf("Failed to update skill: %v", err), http.StatusInternalServerError)
		return
	}

	// Delete old questions
	_, err = tx.Exec("DELETE FROM questions WHERE skill_id = $1", skillId)
	if err != nil {
		_ = tx.Rollback()
		http.Error(w, fmt.Sprintf("Failed to delete old questions: %v", err), http.StatusInternalServerError)
		return
	}

	// Insert new questions
	stmt, err := tx.Prepare("INSERT INTO questions (skill_id, question, correct_answers) VALUES ($1, $2, $3)")
	if err != nil {
		_ = tx.Rollback()
		http.Error(w, fmt.Sprintf("Failed to prepare question statement: %v", err), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	for _, q := range *skill.Questions {
		_, err = stmt.Exec(skillId, q.Question, pq.Array(q.CorrectAnswers))
		if err != nil {
			_ = tx.Rollback()
			http.Error(w, fmt.Sprintf("Failed to insert question: %v", err), http.StatusInternalServerError)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, fmt.Sprintf("Failed to commit transaction: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"skill_id": skillId.String()})
	if err != nil {
		http.Error(w, "Error scanning category", http.StatusInternalServerError)
		return
	}
}

// DeleteSkill (DELETE /skills/{id})
func (t *OpenTutor) DeleteSkill(w http.ResponseWriter, r *http.Request, skillId openapi_types.UUID) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo.UserID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to start transaction: %v", err), http.StatusInternalServerError)
		return
	}

	// Delete associated records first to maintain referential integrity
	_, err = tx.Exec("DELETE FROM tutor_skills WHERE skill_id = $1", skillId)
	if err != nil {
		_ = tx.Rollback()
		http.Error(w, fmt.Sprintf("Failed to delete tutor skills: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("DELETE FROM questions WHERE skill_id = $1", skillId)
	if err != nil {
		_ = tx.Rollback()
		http.Error(w, fmt.Sprintf("Failed to delete skill questions: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("DELETE FROM available_skills WHERE id = $1", skillId)
	if err != nil {
		_ = tx.Rollback()
		http.Error(w, fmt.Sprintf("Failed to delete skill: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, fmt.Sprintf("Failed to commit transaction: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetSkillQuiz retrieves a quiz attempt for a given skill.
func (t *OpenTutor) GetSkillQuiz(w http.ResponseWriter, r *http.Request, skillId openapi_types.UUID) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo.UserID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Check if the tutor already has this skill
	var hasSkill bool
	err := db.GetDB().QueryRow(
		"SELECT EXISTS(SELECT 1 FROM tutor_skills WHERE tutor_id = $1 AND skill_id = $2)",
		authInfo.UserID, skillId,
	).Scan(&hasSkill)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}

	if hasSkill {
		http.Error(w, "Tutor already has this skill.", http.StatusConflict)
		return
	}

	// Create a new quiz attempt
	var attemptID uuid.UUID
	err = db.GetDB().QueryRow(
		"INSERT INTO quiz_attempts (tutor_id, skill_id) VALUES ($1, $2) RETURNING id",
		authInfo.UserID, skillId,
	).Scan(&attemptID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create quiz attempt: %v", err), http.StatusInternalServerError)
		return
	}

	// Select a random set of questions for this skill
	rows, err := db.GetDB().Query(
		fmt.Sprintf("SELECT id, question FROM questions WHERE skill_id = $1 ORDER BY RANDOM() LIMIT %v", N_QUIZ_QUESTIONS),
		skillId,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve questions: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Store the questions in the quiz_attempt_questions table
	quiz := &Quiz{
		Questions: new([]Question),
	}
	for rows.Next() {
		var q Question
		err := rows.Scan(&q.Id, &q.Question)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning questions: %v", err), http.StatusInternalServerError)
			return
		}

		*quiz.Questions = append(*quiz.Questions, q)

		_, err = db.GetDB().Exec(
			"INSERT INTO quiz_attempt_questions (quiz_attempt_id, question_id) VALUES ($1, $2)",
			attemptID, q.Id,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to associate question with attempt: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// Build the response
	quizAttempt := QuizAttempt{
		AttemptId: attemptID,
		SkillId:   skillId,
		Quiz:      *quiz,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(quizAttempt)
	if err != nil {
		http.Error(w, "Error scanning category", http.StatusInternalServerError)
		return
	}
}

func (t *OpenTutor) SubmitSkillQuiz(w http.ResponseWriter, r *http.Request, attemptId openapi_types.UUID) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo.UserID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Decode request body into QuizAttempt struct
	var attempt QuizAttempt
	if err := json.NewDecoder(r.Body).Decode(&attempt); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate that the attempt belongs to the tutor
	var skillId uuid.UUID
	err := db.GetDB().QueryRow(
		"SELECT skill_id FROM quiz_attempts WHERE id = $1 AND tutor_id = $2",
		attemptId, authInfo.UserID,
	).Scan(&skillId)

	if err == sql.ErrNoRows {
		http.Error(w, "Quiz attempt not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch quiz attempt: %v", err), http.StatusInternalServerError)
		return
	}

	// Retrieve the correct answers for the questions in this attempt
	rows, err := db.GetDB().Query(
		"SELECT q.id, q.correct_answers FROM quiz_attempt_questions qa JOIN questions q ON qa.question_id = q.id WHERE qa.quiz_attempt_id = $1",
		attemptId,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Validate and score answers
	var totalScore float64
	var totalQuestions int

	for rows.Next() {
		var questionID uuid.UUID
		var correctAnswers []string
		if err := rows.Scan(&questionID, (*pq.StringArray)(&correctAnswers)); err != nil {
			http.Error(w, fmt.Sprintf("Failed to retrieve question data: %v", err), http.StatusInternalServerError)
			return
		}
		totalQuestions++

		if attempt.Quiz.Questions == nil {
			http.Error(w, "Invalid request: quiz questions are missing", http.StatusBadRequest)
			return
		}

		if len(*attempt.Quiz.Questions) == 0 {
			http.Error(w, "Invalid request: quiz must contain at least one question", http.StatusBadRequest)
			return
		}

		// Find the submitted answer for this question
		var submittedAnswers []string
		for _, q := range *attempt.Quiz.Questions {
			if q.Id == questionID {
				submittedAnswers = q.Answers
				break
			}
		}

		// Use a map to check correct answers efficiently
		correctSet := make(map[string]struct{}, len(correctAnswers))
		for _, correct := range correctAnswers {
			correctSet[correct] = struct{}{}
		}

		// Calculate the fraction of correct answers selected
		if len(correctAnswers) > 0 {
			var correctCount float64
			for _, submitted := range submittedAnswers {
				if _, exists := correctSet[submitted]; exists {
					correctCount++
				}
			}
			// Partial scoring: correct answers selected / total correct answers
			totalScore += correctCount / float64(len(correctAnswers))
		}
	}

	// Calculate final score percentage
	grade := (totalScore / float64(totalQuestions)) * 100
	passed := grade >= PASSING_RATIO*100

	// Update quiz attempt status
	_, err = db.GetDB().Exec(
		"UPDATE quiz_attempts SET passed = $1, completed_at = NOW() WHERE id = $2",
		passed, attemptId,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update quiz attempt: %v", err), http.StatusInternalServerError)
		return
	}

	// If passed, assign the skill
	if passed {
		_, err := db.GetDB().Exec("INSERT INTO tutor_skills (tutor_id, skill_id) VALUES ($1, $2)", authInfo.UserID, skillId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to assign skill: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"grade":  grade,
		"passed": passed,
	})
}

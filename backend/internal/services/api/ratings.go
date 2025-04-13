package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"open-tutor/internal/services/db"
	"open-tutor/middleware"

	"github.com/google/uuid"
)

type RatingRequest struct {
	MeetingId string       `json:"meetingId"`
	Scores    RatingScores `json:"scores"`
	Comment   string       `json:"comment,omitempty"`
}

func (t *OpenTutor) PostRating(w http.ResponseWriter, r *http.Request) {
	// Authenticate user //
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo.UserID == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Parse JSON body //
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var ratingBody RatingRequest
	if err := json.Unmarshal(body, &ratingBody); err != nil {
		http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if user has access to submit rating for this meeting //
	var tutorId string
	err = db.GetDB().QueryRow(`
		SELECT tutor_id FROM meetings WHERE id=$1 AND student_id=$2 LIMIT 1
	`,
		ratingBody.MeetingId,
		authInfo.UserID,
	).Scan(&tutorId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sendError(w, http.StatusBadRequest, fmt.Sprintf("invalid meeting: %v", err))
			return
		}

		err := err.Error()
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("unexpected error: %v", err))
		return
	}

	// Add rating to DB //
	ratingId := uuid.New().String()
	_, err = db.GetDB().Exec(`
		INSERT INTO ratings (id, rating_type, user_id, reviewer_user_id, meeting_id, professionalism, knowledge, communication, punctuality, overall, comment)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
	`,
		ratingId,
		"tutor",
		tutorId,
		authInfo.UserID,
		ratingBody.MeetingId,
		ratingBody.Scores.Professionalism,
		ratingBody.Scores.Knowledge,
		ratingBody.Scores.Communication,
		ratingBody.Scores.Punctuality,
		ratingBody.Scores.Overall,
		ratingBody.Comment,
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			msg := "rating already submitted for this meeting"
			sendError(w, http.StatusBadRequest, "rating already submitted for this meeting")
		} else {
			fmt.Printf("error submitting rating: %v\n", err)
			sendError(w, http.StatusInternalServerError, fmt.Sprintf("error submitting rating: %v", err))
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

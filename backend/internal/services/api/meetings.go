package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"open-tutor/internal/services/db"
	middleware "open-tutor/middleware"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (t *OpenTutor) CreateMeeting(w http.ResponseWriter, r *http.Request) {
	var meetingPayload CreateMeetingBody
	if err := json.NewDecoder(r.Body).Decode(&meetingPayload); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for meeting")
		return
	}

	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo == nil {
		fmt.Printf("No authentication info found.\n")
	} else {
		fmt.Printf("Authenticated user id: %s\n", authInfo.UserID)
	}

	_, err := db.GetDB().Exec(`
		INSERT INTO meetings
		(id, tutor_id, student_id, start_at, end_at, zoom_link)
		VALUES ($1, $2, $3, $4, $5, $6);
	`,
		uuid.New().String(),
		authInfo.UserID,
		meetingPayload.StudentId,
		meetingPayload.StartAt,
		meetingPayload.EndAt,
	)
	if err != nil {
		fmt.Printf("failed to insert meeting into database: %v\n", err)
		sendError(w, http.StatusInternalServerError, "Failed to create meeting")
		return
	}
}

func (t *OpenTutor) DeleteMeetingById(w http.ResponseWriter, r *http.Request, meetingId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) GetMeetingById(w http.ResponseWriter, r *http.Request, meetingId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) UpdateMeetingById(w http.ResponseWriter, r *http.Request, meetingId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

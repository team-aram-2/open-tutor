package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"open-tutor/internal/services/db"
	"open-tutor/middleware"
	"open-tutor/stripe_client"
	"open-tutor/zoom"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stripe/stripe-go/v81"
)

func (t *OpenTutor) CreateMeeting(w http.ResponseWriter, r *http.Request) {
	var meetingPayload CreateMeetingBody

	if err := json.NewDecoder(r.Body).Decode(&meetingPayload); err != nil {
		fmt.Printf("error parsing CreateMeeting body: %s\n", err)
		sendError(w, http.StatusBadRequest, "Invalid format for meeting")
		return
	}

	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo == nil {
		sendError(w, http.StatusUnauthorized, "User is not logged in")
		return
	}

	tutorUserId := authInfo.UserID
	studentUserId := meetingPayload.StudentId

	zoomMeeting, err := zoom.CreateMeeting(tutorUserId, studentUserId.String())
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("error creating Zoom meeting: %s", err))
		return
	}

	_, err = db.GetDB().Exec(`
		INSERT INTO meetings
		(id, tutor_id, student_id, start_at, end_at, zoom_join_link, zoom_host_link)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`,
		uuid.New().String(),
		tutorUserId,
		studentUserId,
		meetingPayload.StartAt,
		meetingPayload.EndAt,
		zoomMeeting.JoinUrl,
		zoomMeeting.StartUrl,
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

func (t *OpenTutor) GetMeetings(w http.ResponseWriter, r *http.Request) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo == nil {
		sendError(w, http.StatusUnauthorized, "User is not logged in")
		return
	}
	userId := authInfo.UserID

	meetingRows, err := db.GetDB().Query(`
		SELECT
			m.id,
			m.tutor_id,
			m.student_id,
			m.start_at,
			m.end_at,
			m.zoom_join_link,
			CASE
        WHEN m.tutor_id = $1
        THEN m.zoom_host_link
        ELSE NULL
    	END AS zoom_host_link
		FROM meetings m
		WHERE
			CASE
				WHEN m.tutor_id = $1
				THEN m.tutor_id = $1
				ELSE m.student_id = $1
			END;
	`, userId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("failed to fetch user role: %s", err))
		return
	}

	meetings := []Meeting{}
	for meetingRows.Next() {
		var meeting Meeting
		err := meetingRows.Scan(
			&meeting.Id,
			&meeting.TutorId,
			&meeting.StudentId,
			&meeting.StartAt,
			&meeting.EndAt,
			&meeting.ZoomJoinLink,
			&meeting.ZoomHostLink,
		)
		if err != nil {
			sendError(w, http.StatusInternalServerError, fmt.Sprintf("failed to parse meetings from database: %s", err))
			return
		}
		meetings = append(meetings, meeting)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(meetings); err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("failed to encode meetings response: %s", err))
		return
	}
}

func (t *OpenTutor) FinalizeMeeting(w http.ResponseWriter, r *http.Request, meetingId uuid.UUID) {
	// Auth check //
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo == nil {
		sendError(w, http.StatusUnauthorized, "User is not logged in")
		return
	}

	// Verify meeting ownership & get hourly rate //
	var tutorHourlyRate int64
	var tutorAccountId string
	var studentUserId string
	var studentCustomerId string
	err := db.GetDB().QueryRow("SELECT tutors.hourly_rate, tutors.stripe_account_id, meetings.student_id FROM meetings LEFT JOIN tutors ON tutors.user_id = meetings.tutor_id WHERE meetings.id = $1 AND meetings.tutor_id = $2", meetingId, authInfo.UserID).Scan(&tutorHourlyRate, &tutorAccountId, &studentUserId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("failed to fetch meeting: %v", err))
		return
	}
	err = db.GetDB().QueryRow("SELECT users.stripe_customer_id FROM users WHERE user_id = $1", studentUserId).Scan(&studentCustomerId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("failed to fetch student customer ID: %v", err))
		return
	}

	// Submit PaymentIntent from student customer to tutor account via Stripe //
	_, err = stripe_client.GetClient().PaymentIntents.New(&stripe.PaymentIntentParams{
		Customer: stripe.String(studentCustomerId),
		TransferData: &stripe.PaymentIntentTransferDataParams{
			Destination: stripe.String(tutorAccountId),
		},
		Amount:   stripe.Int64(tutorHourlyRate * 100), // multiply by 100 for cents value
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
		// Confirm:    stripe.Bool(true),
		// OffSession: stripe.Bool(true),
	})
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("failed to create PaymentIntent for meeting: %v", err))
	}
}

func (t *OpenTutor) UpdateMeetingById(w http.ResponseWriter, r *http.Request, meetingId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

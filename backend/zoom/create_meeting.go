package zoom

import (
	"fmt"
	"time"

	"open-tutor/internal/services/db"
)

type MeetingInvitee struct {
	Email string `json:"email"`
}
type CreateMeetingPayload struct {
	Topic     string    `json:"topic"`
	Type      int       `json:"type"` // 2 for scheduled meeting
	StartTime time.Time `json:"start_time"`
	Duration  int       `json:"duration"` // in minutes
	Settings  struct {
		MeetingInvitees  []MeetingInvitee `json:"meeting_invitees"`
		ParticipantVideo bool             `json:"participant_video"`
		JoinBeforeHost   bool             `json:"join_before_host"`
	} `json:"settings"`
}
type CreateMeetingResponse struct {
	Id       uint64 `json:"id"`
	JoinUrl  string `json:"join_url"`
	StartUrl string `json:"start_url"`
}

func CreateMeeting(tutorId string, studentId string) (*CreateMeetingResponse, error) {
	userRows, err := db.GetDB().Query("SELECT user_id, first_name, last_name, email FROM users WHERE user_id IN ($1, $2)", tutorId, studentId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info from database: %v", err)
	}

	invitees := []MeetingInvitee{}
	var (
		userId    string
		firstName string
		lastName  string
		email     string
	)
	var tutorName string
	for userRows.Next() {
		err = userRows.Scan(&userId, &firstName, &lastName, &email)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %s", err)
		}

		invitees = append(invitees, MeetingInvitee{
			Email: email,
		})
		if userId == tutorId {
			tutorName = firstName + " " + lastName
		}
	}

	meeting := CreateMeetingPayload{
		Topic:     fmt.Sprintf("%s's OpenTutor meeting", tutorName),
		Type:      2,
		StartTime: time.Now().Add(24 * time.Hour),
		Duration:  60,
	}
	meeting.Settings.ParticipantVideo = true
	meeting.Settings.JoinBeforeHost = true
	meeting.Settings.MeetingInvitees = invitees

	var responseBody *CreateMeetingResponse
	responseBody, err = zoomApiRequest[CreateMeetingResponse]("/v2/users/me/meetings", "POST", meeting, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make Zoom API reqeust: %s", err)
	}

	return responseBody, nil
}

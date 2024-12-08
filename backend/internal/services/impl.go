package api

import (
	"encoding/json"
	"net/http"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type OpenTutor struct {
	Users map[string]User
}

var _ ServerInterface = (*OpenTutor)(nil)

func NewOpenTutor() *OpenTutor {
	return &OpenTutor{
		Users: make(map[string]User),
	}
}

func sendError(w http.ResponseWriter, code int, message string) {
	sendError := ErrorModel{
		Code:    code,
		Message: message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(sendError)
}

func (t *OpenTutor) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for user")
		return
	}
	t.Users[string(user.UserId.String())] = user
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

func (t *OpenTutor) CreateMeeting(w http.ResponseWriter, r *http.Request) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
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

func (t *OpenTutor) CreateMessage(w http.ResponseWriter, r *http.Request) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) DeleteMessageById(w http.ResponseWriter, r *http.Request, messageId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) GetMessageById(w http.ResponseWriter, r *http.Request, messageId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) UpdateMessageById(w http.ResponseWriter, r *http.Request, messageId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) CreateMessageAttachment(w http.ResponseWriter, r *http.Request) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) DeleteMessageAttachmentById(w http.ResponseWriter, r *http.Request, messageAttachmentId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) GetMessageAttachmentById(w http.ResponseWriter, r *http.Request, messageAttachmentId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) PostRating(w http.ResponseWriter, r *http.Request) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) SignUpAsStudent(w http.ResponseWriter, r *http.Request) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) GetStudentByID(w http.ResponseWriter, r *http.Request, studentID openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) SignUpAsTutor(w http.ResponseWriter, r *http.Request) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) GetTutorById(w http.ResponseWriter, r *http.Request, tutorId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) DeleteUserById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) GetUserById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) UpdateUserById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

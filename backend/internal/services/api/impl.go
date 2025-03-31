package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"open-tutor/internal/services/db"
	"open-tutor/setup"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type OpenTutor struct{}

var _ ServerInterface = (*OpenTutor)(nil)

func Init() *OpenTutor {
	// Ensure the database is initialized
	_, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	setup.EnsureDefaultAdmin()

	fmt.Println("Database initialized successfully")
	return &OpenTutor{}
}

func sendError(w http.ResponseWriter, code int, message string) {
	sendError := ErrorModel{
		Code:    code,
		Message: message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(sendError)
	fmt.Printf("error on web request %+v\n", sendError)
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

func (t *OpenTutor) GetRatingById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID, params GetRatingByIdParams) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

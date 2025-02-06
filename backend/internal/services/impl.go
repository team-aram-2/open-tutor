package api

import (
	"encoding/json"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
	"fmt"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type OpenTutor struct {}

var _ ServerInterface = (*OpenTutor)(nil)
var db *sql.DB

func NewOpenTutor() *OpenTutor {
	// Connect to database //
	var conErr error
	db, conErr = sql.Open("postgres", "postgres://postgres:developer@db:5432/postgres?sslmode=disable")
	if conErr != nil {
		log.Fatal(conErr)
	}
	// defer db.Close()

	// Test connection //
	conErr = db.Ping()
	if conErr != nil {
		log.Fatal(conErr)
	}

	// Set connection pool settings //
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)
	db.SetConnMaxLifetime(5 * time.Minute)
	fmt.Printf("Connected to database\n")

	return &OpenTutor{}
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

	var userId string
	insertErr := db.QueryRow(`
		INSERT INTO users (email, first_name, last_name)
		VALUES ($1, $2, $3)
		RETURNING user_id
	`,
		user.Email,
		user.FirstName,
		user.LastName,
	).Scan(&userId)

	if insertErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", insertErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"userId": userId,
	})
}

func (t *OpenTutor) CreateMeeting(w http.ResponseWriter, r *http.Request) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) DeleteMeetingById(w http.ResponseWriter, r *http.Request, meetingId interface{}) {
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

func (t *OpenTutor) GetRatingById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID, params GetRatingByIdParams) {
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
	user := &User{}
	selectErr := db.QueryRow(`
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
	json.NewEncoder(w).Encode(user)
}

func (t *OpenTutor) UpdateUserById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

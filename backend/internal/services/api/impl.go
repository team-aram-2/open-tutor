package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"open-tutor/internal/services/db"

	"github.com/google/uuid"
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
}

func checkUser(userId openapi_types.UUID) (bool, error) {
	var exist bool
	err := db.GetDB().QueryRow(`
		SELECT EXISTS (
			SELECT 1 
			FROM users
			WHERE user_id = $1
		)
	`, userId).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (t *OpenTutor) CreateUser(w http.ResponseWriter, r *http.Request) {
	var protoUser ProtoUser
	var user User
	if err := json.NewDecoder(r.Body).Decode(&protoUser); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for user")
		return
	}

	userId := uuid.New().String()
	insertErr := db.GetDB().QueryRow(`
		INSERT INTO users (user_id, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id, email, first_name, last_name;
	`,
		userId,
		protoUser.Email,
		protoUser.FirstName,
		protoUser.LastName,
	).Scan(
		&user.UserId,
		&user.Email,
		&user.FirstName,
		&user.LastName,
	)

	if insertErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", insertErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
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
	var exist bool
	var err error
	exist, err = checkUser(userId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if !exist {
		sendError(w, http.StatusNotFound, "User not found")
		return
	}

	_, deleteErr := db.GetDB().Exec(`
		DELETE FROM users
		WHERE user_id = $1
		`, userId)

	if deleteErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s\n", deleteErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (t *OpenTutor) GetUserById(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID) {
	var exist bool
	var err error
	exist, err = checkUser(userId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if !exist {
		sendError(w, http.StatusNotFound, "User not found")
		return
	}

	user := &User{}
	selectErr := db.GetDB().QueryRow(`
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
	var exist bool
	var err error
	exist, err = checkUser(userId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if !exist {
		sendError(w, http.StatusNotFound, "User not found")
		return
	}

	var user User
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for user")
		return
	}

	updateErr := db.GetDB().QueryRow(`
		UPDATE users
		SET first_name = $1, last_name = $2, email = $3
		WHERE user_id = $4
		RETURNING user_id, email, first_name, last_name;
	`, user.FirstName, user.LastName, user.Email, userId).Scan(
		&updatedUser.UserId,
		&updatedUser.Email,
		&updatedUser.FirstName,
		&updatedUser.LastName,
	)
	if updateErr != nil {
		sendError(w, http.StatusInternalServerError, "Database error")
		fmt.Fprintf(w, "%s\n", updateErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

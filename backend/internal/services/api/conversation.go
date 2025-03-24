package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"open-tutor/internal/services/db"
	"open-tutor/middleware"

	"github.com/google/uuid"
	"github.com/lib/pq"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func checkConversation(conversationId openapi_types.UUID) (bool, error) {
	var exist bool
	err := db.GetDB().QueryRow(`
		SELECT EXISTS (
			SELECT 1
			FROM conversations
			WHERE id = $1
		)
	`, conversationId).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (t *OpenTutor) CreateConversation(w http.ResponseWriter, r *http.Request) {
	var users []openapi_types.UUID
	var convo Conversation
	if parseErr := json.NewDecoder(r.Body).Decode(&users); parseErr != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for conversation")
		return
	}
	conversationId := uuid.New().String()

	insertErr := db.GetDB().QueryRow(`
	INSERT INTO conversations (id, user_ids)
	VALUES ($1, $2)
	RETURNING id;
	`,
		conversationId,
		pq.Array(users),
	).Scan(
		&convo.Id,
	)
	convo.Users = &users

	if insertErr != nil {
		sendError(w, http.StatusInternalServerError, "Database Error")
		fmt.Printf("error creating conversation: %v\n", insertErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(convo)
}

func (t *OpenTutor) GetMessagesByConversationId(w http.ResponseWriter, r *http.Request, conversationId openapi_types.UUID) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo.UserID == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	userid := authInfo.UserID
	var exists bool
	err := db.GetDB().QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)", userid).Scan(&exists)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Database error: %s\n", err)
		return
	}

	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User with ID:{} does not exist\n")
		return
	}
	var exist bool
	var existErr error
	exist, existErr = checkConversation(conversationId)
	if existErr != nil {
		sendError(w, http.StatusInternalServerError, "Database Error (Exist)")
		return
	}
	if !exist {
		sendError(w, http.StatusNotFound, "Conversation not found")
		return
	}

	var access bool
	err = db.GetDB().QueryRow("SELECT EXISTS(SELECT 1 FROM conversations where id = $1 AND $2 = ANY(user_ids))", conversationId, authInfo.UserID).Scan(&access)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Database error: %s\n", err)
		return
	}
	if !access {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
		return
	}

	rows, err := db.GetDB().Query(`
	SELECT *
	FROM messages
	WHERE conversation_id = $1
	`, conversationId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database Error")
		fmt.Printf("error extracting conversation: %v\n", err)
		return
	}
	defer rows.Close()
	var messages []Message

	for rows.Next() {
		var temp Message
		rowErr := rows.Scan(
			&temp.Id,
			&temp.SentOn,
			&temp.OriginId,
			&temp.ConversationId,
			&temp.Message,
		)
		if rowErr != nil {
			sendError(w, http.StatusInternalServerError, "Server error")
			fmt.Printf("error extracting messages: %v\n", err)
			return
		}
		messages = append(messages, temp)
	}

	response := struct {
		Messages []Message `json:"messages"`
	}{
		Messages: messages,
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (t *OpenTutor) GetUsersByConversationId(w http.ResponseWriter, r *http.Request, conversationId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) GetConversationsByUserId(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID) {
	authInfo := middleware.GetAuthenticationInfo(r)
	if authInfo.UserID == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	userid := authInfo.UserID
	var exists bool
	err := db.GetDB().QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)", userid).Scan(&exists)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Database error: %s\n", err)
		return
	}

	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User with ID:{} does not exist\n")
		return
	}
	if userId.String() != userid {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	rows, err := db.GetDB().Query(`
		SELECT id
		FROM conversations
		WHERE $1 = ANY(user_ids)
	`, userId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database Error")
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer rows.Close()

	var conversations []string

	for rows.Next() {
		var temp string
		if err := rows.Scan(&temp); err != nil {
			sendError(w, http.StatusInternalServerError, "Server error")
			return
		}
		conversations = append(conversations, temp)
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(conversations)
}

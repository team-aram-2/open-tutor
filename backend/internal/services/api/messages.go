package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"open-tutor/internal/services/db"
	"open-tutor/middleware"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func checkMessage(messageId openapi_types.UUID) (bool, error) {
	var exist bool
	err := db.GetDB().QueryRow(`
		SELECT EXISTS (
			SELECT 1
			FROM messages
			WHERE id = $1
		)
	`, messageId).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (t *OpenTutor) CreateMessage(w http.ResponseWriter, r *http.Request) {
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
	var proto ProtoMessage
	var message Message
	if parseErr := json.NewDecoder(r.Body).Decode(&proto); parseErr != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for message")
		return
	}
	var access bool
	err = db.GetDB().QueryRow("SELECT EXISTS(SELECT 1 FROM conversations where id = $1 AND $2 = ANY(user_ids))", proto.ConversationId, authInfo.UserID).Scan(&access)
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
	messageId := uuid.New().String()
	timestamp := time.Now()
	insertErr := db.GetDB().QueryRow(`
	INSERT INTO messages (id, sent_at, origin_id, conversation_id, message)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, sent_at, origin_id, conversation_id, message;
	`, messageId, timestamp, proto.OriginId, proto.ConversationId, proto.Message).Scan(
		&message.Id,
		&message.SentOn,
		&message.OriginId,
		&message.ConversationId,
		&message.Message,
	)

	if insertErr != nil {
		sendError(w, http.StatusInternalServerError, "Database error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(message)
}

func (t *OpenTutor) DeleteMessageById(w http.ResponseWriter, r *http.Request, messageId openapi_types.UUID) {
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
	var access bool
	err = db.GetDB().QueryRow("SELECT EXISTS(SELECT 1 FROM messages where id = $1 AND origin_id = $2)", messageId, authInfo.UserID).Scan(&access)
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
	var exist bool
	exist, err = checkMessage(messageId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database Error")
		return
	}
	if !exist {
		sendError(w, http.StatusNotFound, "Message not found")
		return
	}
	_, deleteErr := db.GetDB().Exec(`
		DELETE FROM messages
		WHERE id = $1
	`, messageId)

	if deleteErr != nil {
		sendError(w, http.StatusInternalServerError, "Database Error")
	}

	w.WriteHeader(http.StatusNoContent)
}

func (t *OpenTutor) GetMessageById(w http.ResponseWriter, r *http.Request, messageId openapi_types.UUID) {
	var exist bool
	var err error
	exist, err = checkMessage(messageId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database Error")
		return
	}
	if !exist {
		sendError(w, http.StatusNotFound, "Message not found")
		return
	}
	message := &Message{}
	selectErr := db.GetDB().QueryRow(`
	SELECT *
	FROM messages
	WHERE id = $1
	`, messageId).Scan(
		&message.Id,
		&message.SentOn,
		&message.OriginId,
		&message.ConversationId,
		&message.Message,
	)

	if selectErr != nil {
		sendError(w, http.StatusInternalServerError, "Database Error")
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(message)
}

func (t *OpenTutor) UpdateMessageById(w http.ResponseWriter, r *http.Request, messageId openapi_types.UUID) {
	var exist bool
	var err error
	exist, err = checkMessage(messageId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database Error")
		return
	}
	if !exist {
		sendError(w, http.StatusNotFound, "Message not found")
		return
	}
	var message Message
	var updatedMessage Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		fmt.Printf("Error: %v\n", err)
		sendError(w, http.StatusBadRequest, "Invalid format")
		return
	}

	updateErr := db.GetDB().QueryRow(`
	UPDATE messages
	SET message = $1
	WHERE id = $2
	RETURNING sent_at, origin_id, conversation_id, message;
	`, message.Message, messageId).Scan(
		&updatedMessage.SentOn,
		&updatedMessage.OriginId,
		&updatedMessage.ConversationId,
		&updatedMessage.Message,
	)
	updatedMessage.Id = message.Id

	if updateErr != nil {
		sendError(w, http.StatusInternalServerError, "Database error")
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(updatedMessage)
}

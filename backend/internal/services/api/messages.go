package api

import (
	"encoding/json"
	"net/http"
	"time"

	"open-tutor/internal/services/db"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (t *OpenTutor) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var proto ProtoMessage
	var message Message
	if parseErr := json.NewDecoder(r.Body).Decode(&proto); parseErr != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for message")
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
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) GetMessageById(w http.ResponseWriter, r *http.Request, messageId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) UpdateMessageById(w http.ResponseWriter, r *http.Request, messageId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

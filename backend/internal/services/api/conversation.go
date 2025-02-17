package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"open-tutor/internal/services/db"

	"github.com/google/uuid"
	"github.com/lib/pq"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

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

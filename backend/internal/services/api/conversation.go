package api

import (
	"database/sql"
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
	var users []openapi_types.UUID
	var convo Conversation
	if parseErr := json.NewDecoder(r.Body).Decode(&users); parseErr != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for conversation")
		return
	}
	isOne := false
	for _, user := range users {
		exists, err := checkUser(user)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}
		if !exists {
			sendError(w, http.StatusBadRequest, fmt.Sprintf("User with ID:%s does not exist\n", user))
			return
		}
		if user.String() == userid {
			isOne = true
		}
	}
	if !isOne {
		sendError(w, http.StatusForbidden, "Unauthorized")
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
		SELECT id, user_ids
		FROM conversations
		WHERE $1 = ANY(user_ids)
	`, userId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database Error")
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer rows.Close()

	type convInfo struct {
		id       string
		user_ids []string
	}

	var conversationInfo []convInfo

	for rows.Next() {
		var temp convInfo
		if err := rows.Scan(&temp.id, pq.Array(&temp.user_ids)); err != nil {
			sendError(w, http.StatusInternalServerError, "Server error")
			return
		}
		conversationInfo = append(conversationInfo, temp)
	}

	if err = rows.Err(); err != nil {
		sendError(w, http.StatusInternalServerError, "Server error")
		return
	}

	userMap := make(map[string]string)

	var toReturn []ConversationName

	for _, conv := range conversationInfo {
		var result ConversationName
		result.Id = &conv.id
		var names []string

		for _, uid := range conv.user_ids {
			fullname, exists := userMap[uid]
			if !exists {
				var first, last string

				err := db.GetDB().QueryRow(`
				SELECT first_name, last_name
				FROM users
				WHERE user_id = $1`, uid).Scan(&first, &last)

				if err == sql.ErrNoRows {
					fullname = "Unknown"
				} else if err != nil {
					sendError(w, http.StatusInternalServerError, "Server Error")
					return
				} else {
					fullname = fmt.Sprintf("%s %s", first, last)
				}
				userMap[uid] = fullname
			}
			names = append(names, fullname)
		}
		var namesList string
		for i, name := range names {
			if i > 0 {
				namesList += ", "
			}
			namesList += name
		}
		result.Name = &namesList
		toReturn = append(toReturn, result)
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(toReturn)
}

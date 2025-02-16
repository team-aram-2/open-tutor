package api

import (
	"net/http"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (t *OpenTutor) CreateConversation(w http.ResponseWriter, r *http.Request) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) GetMessagesByConversationId(w http.ResponseWriter, r *http.Request, conversationId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

func (t *OpenTutor) GetUsersByConversationId(w http.ResponseWriter, r *http.Request, conversationId openapi_types.UUID) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

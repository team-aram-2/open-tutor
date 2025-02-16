package api

import (
	"net/http"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

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

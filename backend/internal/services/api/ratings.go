package api

import "net/http"

func (t *OpenTutor) PostRating(w http.ResponseWriter, r *http.Request) {
	sendError(w, http.StatusMethodNotAllowed, "TODO")
}

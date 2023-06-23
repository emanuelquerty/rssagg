package handlers

import (
	"net/http"

	"github.com/emanuelquerty/rssagg/internal/responder"
)

func HandleError(w http.ResponseWriter, r *http.Request) {
	responder.RespondWithError(w, 400, "Something went wrong")
}

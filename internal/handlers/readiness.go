package handlers

import (
	"net/http"

	"github.com/emanuelquerty/rssagg/internal/responder"
)

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	responder.RespondWithJSON(w, 200, struct{}{})
}

package handlers

import (
	"net/http"

	"github.com/emanuelquerty/rssagg/internal/responder"
	"github.com/emanuelquerty/rssagg/internal/services"
)

func (hctx *HandlerContext) GetUser(w http.ResponseWriter, r *http.Request, user services.User) {
	responder.RespondWithJSON(w, 200, user)
}

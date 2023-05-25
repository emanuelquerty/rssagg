package handlers

import (
	"net/http"

	"github.com/emanuelquerty/rssagg/internal/services"
	"github.com/emanuelquerty/rssagg/internal/utils"
)

func (hctx *HandlerContext) GetUser(w http.ResponseWriter, r *http.Request, user services.User) {
	utils.RespondWithJSON(w, 200, user)
}

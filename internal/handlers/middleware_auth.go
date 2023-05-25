package handlers

import (
	"fmt"
	"net/http"

	"github.com/emanuelquerty/rssagg/internal/auth"
	"github.com/emanuelquerty/rssagg/internal/services"
	"github.com/emanuelquerty/rssagg/internal/utils"
)

type authHandler func(http.ResponseWriter, *http.Request, services.User)

func (hctx *HandlerContext) MiddlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, 401, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := services.GetUser(r.Context(), hctx.DBConn, apikey)
		if err != nil {
			hctx.Logger.Println("Auth error:", err)
			utils.RespondWithError(w, 401, fmt.Sprintf("Auth error - Api key is incorrect: %v", err))
			return
		}
		handler(w, r, user)
	}
}

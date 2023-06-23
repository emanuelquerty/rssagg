package handlers

import (
	"net/http"

	"github.com/emanuelquerty/rssagg/internal/responder"
	"github.com/emanuelquerty/rssagg/internal/services"
)

func (hctx *HandlerContext) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := services.GetFeeds(r.Context(), hctx.DBConn)
	if err != nil {
		hctx.Logger.Println(err)
		responder.RespondWithError(w, 500, "No feeds found")
		return
	}
	responder.RespondWithJSON(w, 200, feeds)
}

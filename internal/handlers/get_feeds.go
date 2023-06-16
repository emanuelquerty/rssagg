package handlers

import (
	"net/http"

	"github.com/emanuelquerty/rssagg/internal/services"
	"github.com/emanuelquerty/rssagg/internal/utils"
)

func (hctx *HandlerContext) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := services.GetFeeds(r.Context(), hctx.DBConn)
	if err != nil {
		hctx.Logger.Println(err)
		utils.RespondWithError(w, 500, "No feeds found")
		return
	}
	utils.RespondWithJSON(w, 200, feeds)
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/emanuelquerty/rssagg/internal/services"
	"github.com/emanuelquerty/rssagg/internal/utils"
)

func (hctx *HandlerContext) GetFeedFollows(w http.ResponseWriter, r *http.Request, user services.User) {
	feedFollows, err := services.GetFeedFollows(r.Context(), hctx.DBConn, user.ID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("could not get the feeds followed by the user with given api key %v", err))
		return
	}
	type responsData struct {
		Following []services.FeedFollows `json:"following,omitempty"  bson:"following"`
	}
	utils.RespondWithJSON(w, 200, responsData{
		Following: feedFollows,
	})
}

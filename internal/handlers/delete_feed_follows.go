package handlers

import (
	"fmt"
	"net/http"

	"github.com/emanuelquerty/rssagg/internal/services"
	"github.com/emanuelquerty/rssagg/internal/utils"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (hctx *HandlerContext) DeleteFeedFollows(w http.ResponseWriter, r *http.Request, user services.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprint("Error parsing url params:", err))
		return
	}

	err = services.DeleteFeedFollows(r.Context(), hctx.DBConn, feedFollowID, user.ID)
	if err != nil {
		utils.RespondWithError(w, 400, err.Error())
		return
	}
	type responsData struct {
		Msg string `json:"msg,omitempty"  bson:"msg"`
	}
	utils.RespondWithJSON(w, 200, responsData{
		Msg: "Feed has been unfollowed",
	})
}

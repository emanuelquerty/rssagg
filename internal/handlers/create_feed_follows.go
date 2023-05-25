package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/emanuelquerty/rssagg/internal/services"
	"github.com/emanuelquerty/rssagg/internal/utils"
	"github.com/google/uuid"
)

func (hctx *HandlerContext) CreateFeedFollows(w http.ResponseWriter, r *http.Request, user services.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id,omitempty"  bson:"feed_id"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprint("Error parsing JSON:", err))
		return
	}

	feedFollowsParams := services.FeedFollows{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	}
	err = services.CreateFeedFollows(r.Context(), hctx.DBConn, feedFollowsParams)
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprint("Could not follow the feed:", err))
		return
	}
	type responseData struct {
		Msg       string               `json:"msg,omitempty"  bson:"msg"`
		Following services.FeedFollows `json:"following,omitempty"  bson:"following"`
	}
	utils.RespondWithJSON(w, 201, responseData{
		Msg:       "You are following the given feed now",
		Following: feedFollowsParams,
	})
}

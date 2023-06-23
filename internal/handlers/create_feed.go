package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/emanuelquerty/rssagg/internal/responder"
	"github.com/emanuelquerty/rssagg/internal/services"
	"github.com/google/uuid"
)

func (hctx *HandlerContext) CreateFeed(w http.ResponseWriter, r *http.Request, user services.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url,omitempty"  bson:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responder.RespondWithError(w, 400, fmt.Sprint("Error parsing JSON:", err))
		return
	}

	feed, err := services.CreateFeed(r.Context(), hctx.DBConn, services.Feed{
		ID:            uuid.New(),
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		Name:          params.Name,
		URL:           params.URL,
		UserID:        user.ID,
		LastFetchedAt: time.Now().UTC(),
	})
	if err != nil {
		hctx.Logger.Println("Couldn't create feed", err)
		responder.RespondWithError(w, 400, "Couldn't create feed. url is already associated with an existing feed.")
		return
	}
	responder.RespondWithJSON(w, 201, feed)
}

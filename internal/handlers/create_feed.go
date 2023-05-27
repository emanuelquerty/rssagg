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

func (hctx *HandlerContext) CreateFeed(w http.ResponseWriter, r *http.Request, user services.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url,omitempty"  bson:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprint("Error parsing JSON:", err))
		return
	}

	feed, err := services.CreateFeed(r.Context(), hctx.DBConn, services.Feed{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		URL:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		hctx.Logger.Println("Couldn't create user", err)
		utils.RespondWithError(w, 500, fmt.Sprint("Couldn't create feed:", err))
		return
	}
	utils.RespondWithJSON(w, 201, feed)
}
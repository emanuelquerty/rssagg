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

func (hctx *HandlerContext) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprint("Error parsing JSON:", err))
		return
	}

	user, err := services.CreateUser(r.Context(), hctx.DBConn, services.User{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		hctx.Logger.Println("Couldn't create user", err)
		utils.RespondWithError(w, 500, fmt.Sprint("Couldn't create user:", err))
		return
	}
	utils.RespondWithJSON(w, 201, user)
}

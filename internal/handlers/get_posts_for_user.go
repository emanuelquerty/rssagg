package handlers

import (
	"fmt"
	"net/http"

	"github.com/emanuelquerty/rssagg/internal/services"
	"github.com/emanuelquerty/rssagg/internal/utils"
)

func (hctx *HandlerContext) GetPosts(w http.ResponseWriter, r *http.Request, user services.User) {
	posts, err := services.GetPosts(r.Context(), hctx.DBConn, user.ID, 4)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Could not get posts: %v", err))
		return
	}
	utils.RespondWithJSON(w, 200, posts)
}

package handlers

import (
	"net/http"

	"github.com/emanuelquerty/rssagg/internal/utils"
)

func HandleError(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 400, "Something went wrong")
}

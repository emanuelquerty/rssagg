package handlers

import (
	"net/http"

	"github.com/emanuelquerty/rssagg/internal/utils"
)

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, struct{}{})
}

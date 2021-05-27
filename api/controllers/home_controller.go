package controllers

import (
	"net/http"

	"github.com/kjuchniewicz/fullstack/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Witamy w naszym API")
}

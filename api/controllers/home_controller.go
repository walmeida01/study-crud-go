package controllers

import (
	"net/http"

	"github.com/walmeida01/study-fullstack/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this Awesome API")
}

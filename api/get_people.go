package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_task/pkg"
)

func (s *Server) HandleGetNames(ctx *gin.Context) {
	var people []pkg.Person
	s.DB.Find(&people)

	ctx.JSON(http.StatusOK, people)
}

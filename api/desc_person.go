package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_task/pkg"
)

func (s *Server) SortedPeopleDESC(ctx *gin.Context) {
	var people []pkg.Person
	s.DB.Order("name desc").Find(&people)

	ctx.JSON(http.StatusOK, people)
}

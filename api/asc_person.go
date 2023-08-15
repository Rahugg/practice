package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_task/pkg"
)

func (s *Server) SortedPeopleASC(ctx *gin.Context) {
	var people []pkg.Person
	s.DB.Order("name").Find(&people)

	ctx.JSON(http.StatusOK, people)
}

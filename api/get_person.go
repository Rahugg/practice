package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"test_task/pkg"
)

func (s *Server) HandleGetByName(ctx *gin.Context) {
	var person pkg.Person

	name := ctx.Param("name")

	ret := s.DB.First(&person, "name = ?", name)

	if ret.RowsAffected == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if ret.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, ret.Error)
		return
	}

	ctx.JSON(http.StatusOK, person)
}

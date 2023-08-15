package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
	"log"
	"net/http"
	"test_task/pkg"
)

func (s *Server) HandleAddPerson(ctx *gin.Context) {
	var person pkg.Person

	err := ctx.BindJSON(&person)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	person.Id = ulid.Now()

	r := s.DB.Create(&person)

	if r.Error != nil {
		log.Fatal(r.Error)
	}

	s.DB.Save(&person)

	ctx.JSON(http.StatusOK, &person)
}

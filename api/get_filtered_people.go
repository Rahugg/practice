package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test_task/pkg"
)

func (s *Server) HandleFilterByAge(ctx *gin.Context) {
	minAgeStr := ctx.DefaultQuery("min_age", "0")
	maxAgeStr := ctx.DefaultQuery("max_age", "99")

	var minAge, maxAge int
	var err error

	// Parse minAge and maxAge parameters
	if minAge, err = strconv.Atoi(minAgeStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid min_age parameter",
		})
		return
	}
	if maxAge, err = strconv.Atoi(maxAgeStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid max_age parameter",
		})
		return
	}

	var filteredResults []pkg.Person

	// Query the database for people within the specified age range
	err = s.DB.Where("age >= ? AND age <= ?", minAge, maxAge).Order("age").Find(&filteredResults).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"results": filteredResults,
	})
}

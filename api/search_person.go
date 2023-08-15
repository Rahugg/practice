package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_task/pkg"
)

func (s *Server) HandleSearchPersonByName(ctx *gin.Context) {
	searchTerm := ctx.Query("name")

	if searchTerm == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "search term 'name' is required",
		})
		return
	}

	var searchResults []pkg.Person

	err := s.DB.Where("name LIKE ?", "%"+searchTerm+"%").Find(&searchResults).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	if len(searchResults) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "person not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"results": searchResults,
	})
}
func (s *Server) HandleSearchPersonBySurname(ctx *gin.Context) {
	searchTerm := ctx.Query("surname")

	if searchTerm == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "search term 'surname' is required",
		})
		return
	}

	var searchResults []pkg.Person

	err := s.DB.Where("surname LIKE ?", "%"+searchTerm+"%").Find(&searchResults).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	if len(searchResults) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "person not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"results": searchResults,
	})
}

func (s *Server) HandleSearchPersonByAddress(ctx *gin.Context) {
	searchTerm := ctx.Query("address")

	if searchTerm == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "search term 'address' is required",
		})
		return
	}

	var searchResults []pkg.Person

	err := s.DB.Where("address LIKE ?", "%"+searchTerm+"%").Find(&searchResults).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	if len(searchResults) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "person not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"results": searchResults,
	})
}

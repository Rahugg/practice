package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Server struct {
	DB  *gorm.DB
	Gin *gin.Engine
}

func (s *Server) InitDb(dsn string) *Server {
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal(err)
	}

	s.DB = db

	return s
}

func (s *Server) InitGin() *Server {
	g := gin.Default()

	s.Gin = g

	return s
}

func (s *Server) Ready() bool {
	return s.DB != nil && s.Gin != nil
}

func (s *Server) Start(ep string) error {
	if !s.Ready() {
		return errors.New("server isn't ready - make sure to init db and gin")
	}

	if err := http.ListenAndServe(ep, s.Gin.Handler()); err != nil {
		log.Fatal(err)
	}

	return nil
}

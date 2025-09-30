package server

import (
	"github.com/Molert511/TextScan/internal/delivery/http"
	"github.com/Molert511/TextScan/internal/repository"
	"github.com/Molert511/TextScan/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	engine := gin.Default()

	engine.Static("/", "./web")

	repo := repository.NewOCRRepository()
	uc := usecase.NewOCRUsecase(repo)
	http.NewHandler(engine, uc)

	return &Server{engine: engine}
}

func (s *Server) Run() error {
	return s.engine.Run(":8080")
}

package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyStreamAuthServer struct {
	configure Configure
	engin     *gin.Engine
}

func InitializeServer() *MyStreamAuthServer {
	conf := Configure{StorageType(0)}
	ginEngin := gin.Default()

	return &MyStreamAuthServer{
		configure: conf,
		engin:     ginEngin,
	}
}

func (s *MyStreamAuthServer) Handle() {
	// Ping test
	s.engin.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func (s *MyStreamAuthServer) Run() {
	s.engin.Run(":8000")
}

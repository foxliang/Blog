package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mysql/mysql"
)

// Server api server
type Server struct {
	engine *gin.Engine
}

// NewServer create api server
func NewServer() *Server {
	server := &Server{}
	return server
}

// Start start api server
func (s Server) Start() {
	fmt.Println("api server start", "127.0.0.1:8088")
	gin.SetMode(gin.ReleaseMode)
	s.engine = gin.New()

	s.engine.Use(gin.Recovery())

	// init mysql
	mysql.GetDB()

	// init route
	s.registeRoute()

	err := s.engine.Run(":8088")
	if err != nil {
		fmt.Println("api server stopped", err)
		return
	}
}

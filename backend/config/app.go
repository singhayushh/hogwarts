package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func (server *Server) Init(port string) {
	server.Router = gin.New()
	// gin.SetMode(gin.ReleaseMode)
	server.InitRoutes()
	server.Router.LoadHTMLGlob("frontend/views/**/*")
	server.Router.Static("/assets", "frontend/assets")
	fmt.Println("Listening to port", port)
	server.Router.Run(port)
}

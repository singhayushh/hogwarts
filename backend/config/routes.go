package config

import "github.com/singhayushh/hogwarts/backend/controllers"

func (server *Server) InitRoutes() {
	index := new(controllers.IndexController)
	server.Router.GET("/", index.RenderHome)
}

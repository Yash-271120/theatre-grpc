package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	return &Router{gin.Default()}
}

func (r *Router) Start() {
	host, hostexists := os.LookupEnv("HOST")
	port, portexists := os.LookupEnv("PORT")

	if !hostexists {
		log.Println("host not fount using default - 127.0.0.1")
		host = "127.0.0.1"
	}
	if !portexists {
		log.Println("port not fount using default - 5000")
		port = "5000"
	}
	r.getRoutes()
	r.Engine.Run(host + ":" + port)
}

func (r *Router) getRoutes() {
	v1 := r.Group("/api/v1")

	addMovieRoutes(v1)
	addHallRoutes(v1)
	addTicketRoutes(v1)
}

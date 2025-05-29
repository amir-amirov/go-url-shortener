package server

import (
	"github.com/amir-amirov/go-url-shortener/internal/shorten"
	"github.com/gin-gonic/gin"
)

type Server struct {
	g         *gin.Engine
	shortener *shorten.Service
}

func New(shortener *shorten.Service) *Server {
	server := &Server{
		g:         gin.Default(),
		shortener: shortener}
	server.registerRoutes(server.g)
	return server
}

func (s *Server) registerRoutes(server *gin.Engine) {
	server.GET(":identifier", HandleRedirect(s.shortener)) // redirect to original URL
	server.POST("/shorten", HandleShorten(s.shortener))    // handle shorten
}

func (s *Server) Run(addr string) error {
	return s.g.Run(addr)
}

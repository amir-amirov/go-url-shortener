package server

import (
	"github.com/amir-amirov/go-url-shortener/internal/shorten"
	"github.com/gin-contrib/cors"
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

	server.g.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Allows all origins
		AllowMethods:     []string{"GET", "POST", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	server.registerRoutes(server.g)
	return server
}

func (s *Server) registerRoutes(server *gin.Engine) {

	server.GET(":identifier", HandleRedirect(s.shortener)) // redirect to original URL
	server.POST("/shorten", HandleShorten(s.shortener))    // handle shorten

	server.Static("/assets", "frontend/react-url-shortener/dist/assets")
	// server.Static("/assets", "../frontend/react-url-shortener/dist/assets")

	// Serve index.html for the root endpoint and other unmatched routes (SPA support)
	server.StaticFile("/", "frontend/react-url-shortener/dist/index.html")
	// server.StaticFile("/", "../frontend/react-url-shortener/dist/index.html")

	// Serve frontend
	server.NoRoute(HandleWebRoute2)

}

func (s *Server) Run(addr string) error {
	return s.g.Run(addr)
}

func HandleWebRoute2(c *gin.Context) {
	c.File("../../frontend/react-url-shortener/dist/index.html")
}

package server

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleWebRoute(staticFiles embed.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := staticFiles.ReadFile("frontend/react-url-shortener/dist/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "index.html not found")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", file)
	}
}

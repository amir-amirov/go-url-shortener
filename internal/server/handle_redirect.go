package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/amir-amirov/go-url-shortener/internal/model"
	"github.com/gin-gonic/gin"
)

type redirecter interface {
	Redirect(ctx context.Context, identifier string) (string, error)
}

func HandleRedirect(redirecter redirecter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		identifier := ctx.Param("identifier")
		fmt.Println("Got identifier:", identifier)

		redirectURL, err := redirecter.Redirect(context.TODO(), identifier)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": "URL not found",
				})
				return
			}

			log.Printf("error getting redirect url for %q: %v", identifier, err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, redirectURL)
	}
}

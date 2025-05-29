package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/amir-amirov/go-url-shortener/internal/model"
	"github.com/amir-amirov/go-url-shortener/internal/shorten"
	"github.com/gin-gonic/gin"
	. "github.com/samber/mo"
)

type shortener interface {
	Shorten(context.Context, model.ShortenInput) (*model.Shortening, error)
}

type shortenRequest struct {
	URL        string `json:"url" validate:"required,url"`
	Identifier string `json:"identifier,omitempty" validate:"omitempty,alphanum"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url,omitempty"`
	Message  string `json:"message,omitempty"`
}

func HandleShorten(shortener shortener) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req shortenRequest

		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request format",
			})
		}

		identifier := None[string]()
		if strings.TrimSpace(req.Identifier) != "" {
			identifier = Some(req.Identifier)
		}

		input := model.ShortenInput{
			RawURL:     req.URL,
			Identifier: identifier,
		}

		shortening, err := shortener.Shorten(context.TODO(), input)
		if err != nil {
			if errors.Is(err, model.ErrIdentifierExists) {
				ctx.JSON(http.StatusConflict, gin.H{"message": model.ErrIdentifierExists.Error()})
				return
			}

			log.Printf("error shortening url %q: %v", req.URL, err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error"})
			return
		}

		shortURL, err := shorten.PrependBaseURL("http://localhost:8080", shortening.Identifier)
		if err != nil {
			log.Printf("error generating full url for %q: %v", shortening.Identifier, err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":   "Successfully shortened URL",
			"short_url": shortURL})

	}
}

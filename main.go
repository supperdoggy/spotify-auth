package spotify_auth

import (
	"github.com/gin-gonic/gin"
	handlers2 "github.com/supperdoggy/spotify-web-project/spotify-auth/internal/handlers"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()
	logger, _ := zap.NewDevelopment()
	handlers := handlers2.NewHandlers(logger)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/new_token")
	}

	if err := r.Run(":8083"); err != nil {
		logger.Fatal("error running application")
	}
}

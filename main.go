package main

import (
	"github.com/gin-gonic/gin"
	db2 "github.com/supperdoggy/spotify-web-project/spotify-auth/internal/db"
	handlers2 "github.com/supperdoggy/spotify-web-project/spotify-auth/internal/handlers"
	"github.com/supperdoggy/spotify-web-project/spotify-auth/internal/service"
	"go.uber.org/zap"
)

// TODO TEST!!!!
func main() {
	r := gin.Default()
	logger, _ := zap.NewDevelopment()
	db, err := db2.NewDB("spotify", logger)
	if err != nil {
		logger.Fatal("error connecting to db", zap.Error(err))
	}
	s := service.NewService(logger, db)
	handlers := handlers2.NewHandlers(logger, s)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/new_token", handlers.NewToken)
		apiv1.POST("/check_token", handlers.CheckToken)
		apiv1.POST("/register", handlers.Register)
		apiv1.POST("/login", handlers.Login)
	}

	if err := r.Run(":8083"); err != nil {
		logger.Fatal("error running application")
	}
}

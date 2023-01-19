package pkg

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shankalpa12/config"
)

type HTTPServer struct {
	*gin.Engine
}

func NewHTTPServer(env config.Config) HTTPServer {
	httpRouter := gin.Default()

	httpRouter.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
				AllowHeaders:     []string{"*"},
				AllowCredentials: true,
			},
		),
	)

	httpRouter.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Practice API Up and Running"})
	})

	return HTTPServer{httpRouter}
}

func (sh *HTTPServer) Start(env config.Config) {
	if env.ServerPort == "" {
		if err := sh.Run(); err != nil {
			fmt.Println("ERROR OCCURED DURING THE RUNNING OF APPLICATION")
			return
		}
	}

	if err := sh.Run(":" + env.ServerPort); err != nil {
		fmt.Println("ERROR OCCURED DURING THE RUNNING OF APPLICATION")
		return
	}
}

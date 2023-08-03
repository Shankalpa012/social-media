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

type Login struct {
	NAME  string `json:"name"`
	EMAIL string `json:"email" `
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

	httpRouter.POST("/user", func(ctx *gin.Context) {
		login := Login{}
		err := ctx.Bind(&login)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"data": err.Error()})
			return
		}

		fmt.Println(len(login.NAME))

		if login.NAME == "" || len(login.NAME) > 10 {
			ctx.JSON(http.StatusBadGateway, gin.H{"Data": " Validation Error"})
			return
		}
		fmt.Println(login.EMAIL)
		ctx.JSON(http.StatusOK, gin.H{"dataaaaaa": login.EMAIL})
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

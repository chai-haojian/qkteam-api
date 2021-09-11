package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qkteam-api/controller"
	"qkteam-api/logger"
	"qkteam-api/middleware"
)

func Setup() *gin.Engine {
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.Cors())
	register := r.Group("/register")
	{
		register.POST("/submit", controller.SubmitHandler)
	}
	admin := r.Group("/admin")
	{
		admin.GET("/all", controller.GetAllPostHandler)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "http server ok!",
		})
	})

	return r
}

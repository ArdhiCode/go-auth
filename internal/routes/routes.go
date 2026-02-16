package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {

	auth := r.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}

}

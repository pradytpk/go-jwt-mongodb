package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pradytpk/go-jwt/controllers"
	"github.com/pradytpk/go-jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/user/:user_id", controllers.GetUser())
}

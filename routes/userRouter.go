package routes

import (
	controller "github.com/amardeep/golang-jwt-project/controllers"
	"github.com/amardeep/golang-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}

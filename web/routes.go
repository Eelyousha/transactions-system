// routes.go

package web

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func InitializeRoutes() {

	// Handle the index route
	Router.GET("/", GetUsers)

	Router.GET("/user_state/:nickname/lk", GetCurrentState)

	Router.POST("/user_state/:nickname/addon", AddRequest)
}

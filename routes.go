// routes.go

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", getUsers)

	router.GET("/user_state/:nickname/lk", getCurrentState)

	router.POST("/user_state/:nickname/addon", addRequest)
}

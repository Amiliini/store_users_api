package app

import (
	"kauppa/controllers/ping"
	"kauppa/controllers/users"
)

func mapUrls() {

	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}

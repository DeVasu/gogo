package app

import (
	"gogo/controllers/users"
)

func mapUrls() {

	router.POST("/users", users.Create)
	router.GET("/:user_name", users.GetByName)

}

package main

import (
	"net/http"

	"angel_clothes.make_friends/m/v2/controllers/routes"
	"angel_clothes.make_friends/m/v2/models"
	"angel_clothes.make_friends/m/v2/sql"
)

func init() {
	db := sql.OpenDbConnection()
	// migration
	db.AutoMigrate(&models.User{}, &models.Person{})
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := routes.SetupRouter()
	r.StaticFS("/images", http.Dir("./images"))
	r.Run(":10086") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

package routes

import (
	"database/sql"
	"go_products_pagination/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*/*.html")

	r.Static("/assets", "./assets")

	r.GET("/", controllers.Productlist)

	return r
}

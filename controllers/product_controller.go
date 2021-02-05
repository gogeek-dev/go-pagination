package controllers

import (
	"fmt"
	mysqldb "go_products_pagination/connection"
	"go_products_pagination/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Productlist(c *gin.Context) {

	db := mysqldb.SetupDB()

	lmt, _ := strconv.Atoi(c.Query("limit"))

	limit := 0 + lmt

	ProductsRows, err := db.Query("select id,title,description,price,image_path from tbl_products limit ?,6", limit)

	if err != nil {

		fmt.Println(err)
	}
	products := models.Products{}

	res := []models.Products{}

	for ProductsRows.Next() {

		var id int

		var title, description, image_path string

		var price float32
		err = ProductsRows.Scan(&id, &title, &description, &price, &image_path)
		if err != nil {
			fmt.Println(err)
		}

		products.ID = id

		products.Title = title

		products.Price = price

		products.Discription = description

		products.Imagepath = image_path

		res = append(res, products)

	}

	var pcount int

	db.QueryRow("SELECT count(id)  FROM tbl_products").Scan(&pcount)

	var Previous, Next int
	if limit <= 0 {
		Previous = 0
		Next = limit + 6
	} else if limit < pcount-6 {
		Previous = limit - 6
		Next = limit + 6
	} else {
		Previous = limit - 6
		Next = limit
	}

	c.HTML(200, "products.html", gin.H{"products": res, "Next": Next, "Previous": Previous})
}

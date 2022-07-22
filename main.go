package main

import (
	"log"
	"net/http"

	"github.com/daopmdean/food-delivery-go/modules/restaurant/restauranttransport/ginrestaurant"
	"github.com/daopmdean/food-delivery-go/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=mystrongpassword dbname=food-delivery port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("error init db", err)
	}

	appContext := tool.NewAppContext(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(*appContext))
	}

	r.Run()
}

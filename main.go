package main

import (
	"log"
	"net/http"
	"os"

	"github.com/daopmdean/food-delivery-go/middleware"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restauranttransport/ginrestaurant"
	"github.com/daopmdean/food-delivery-go/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("error init db", err)
	}

	appContext := tool.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(*appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	restaurants := r.Group("/restaurants")
	{
		restaurants.GET("", ginrestaurant.ListRestaurant(*appContext))
		restaurants.POST("", ginrestaurant.CreateRestaurant(*appContext))
		restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(*appContext))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(*appContext))
	}

	r.Run()
}

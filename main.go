package main

import (
	"log"
	"net/http"

	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantstorage"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantusecase"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type appContext struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{
		db: db,
	}
}

func main() {
	dsn := "host=localhost user=postgres password=mystrongpassword dbname=food-delivery port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("error init db", err)
	}

	appContext := NewAppContext(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", func(ctx *gin.Context) {
			var data restaurantmodel.RestaurantCreate

			err := ctx.ShouldBind(&data)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error":   err,
					"message": "can not bind data",
				})
				return
			}

			store := restaurantstorage.NewSqlStore(appContext.db)
			usecase := restaurantusecase.NewCreateRestaurantUsecase(store)

			err = usecase.CreateRestaurant(ctx.Request.Context(), &data)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error":   err,
					"message": "can not create restaurant",
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": "restaurant created",
			})
		})
	}

	r.Run()
}

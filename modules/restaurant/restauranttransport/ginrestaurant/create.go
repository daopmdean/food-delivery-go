package ginrestaurant

import (
	"net/http"

	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantstorage"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantusecase"
	"github.com/daopmdean/food-delivery-go/tool"
	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx tool.AppContext) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		err := ctx.ShouldBind(&data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"message": "can not bind data",
			})
			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetDBConnection())
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
	}
}

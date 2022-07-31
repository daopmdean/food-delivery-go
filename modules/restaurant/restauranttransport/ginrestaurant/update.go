package ginrestaurant

import (
	"net/http"
	"strconv"

	"github.com/daopmdean/food-delivery-go/common"

	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantstorage"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantusecase"
	"github.com/daopmdean/food-delivery-go/tool"
	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(appCtx tool.AppContext) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		var data restaurantmodel.RestaurantUpdate

		err := ctx.ShouldBind(&data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetDBConnection())
		usecase := restaurantusecase.NewUpdateRestaurantUsecase(store)

		err = usecase.UpdateRestaurant(ctx.Request.Context(), id, &data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse("restaurant updated"))
	}
}

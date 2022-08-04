package ginrestaurant

import (
	"net/http"
	"strconv"

	"github.com/daopmdean/food-delivery-go/common"

	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantstorage"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantusecase"
	"github.com/daopmdean/food-delivery-go/tool"
	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx tool.AppContext) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetDBConnection())
		usecase := restaurantusecase.NewDeleteRestaurantUsecase(store)

		err = usecase.SoftDelete(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse("restaurant deleted"))
	}
}

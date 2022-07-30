package ginrestaurant

import (
	"net/http"

	"github.com/daopmdean/food-delivery-go/common"

	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantstorage"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantusecase"
	"github.com/daopmdean/food-delivery-go/tool"
	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx tool.AppContext) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		var filter restaurantmodel.Filter

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		paging.Fill()

		store := restaurantstorage.NewSqlStore(appCtx.GetDBConnection())
		usecase := restaurantusecase.NewListRestaurantUsecase(store)

		result, err := usecase.ListRestaurant(ctx.Request.Context(), &filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

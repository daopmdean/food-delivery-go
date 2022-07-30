package middleware

import (
	"fmt"
	"net/http"

	"github.com/daopmdean/food-delivery-go/common"
	"github.com/daopmdean/food-delivery-go/tool"
	"github.com/gin-gonic/gin"
)

func Recover(appCtx tool.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
				}

				appErr := common.ErrInternal(fmt.Errorf("%v", err)).
					Status(http.StatusInternalServerError)

				ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		ctx.Next()
	}
}

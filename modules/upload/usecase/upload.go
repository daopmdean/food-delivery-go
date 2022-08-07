package uploadusecase

import (
	"context"

	"github.com/daopmdean/food-delivery-go/common"
)

type CreateImageStorage interface {
	CreateImage(ctx context.Context, data *common.Image) error
}

type uploadUsecase struct {
}

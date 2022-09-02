package uploadusecase

import (
	"context"
	"errors"

	"github.com/daopmdean/food-delivery-go/common"
)

type CreateImageStorage interface {
	CreateImage(ctx context.Context, data *common.Image) error
}

type uploadUsecase struct {
	storage CreateImageStorage
}

func NewUploadUsecase(storage CreateImageStorage) *uploadUsecase {
	return &uploadUsecase{storage: storage}
}

func (usecase *uploadUsecase) Upload(
	ctx context.Context,
	data []byte,
	folder, filename string,
) (*common.Image, error) {
	return nil, errors.New("not implement yet")
}

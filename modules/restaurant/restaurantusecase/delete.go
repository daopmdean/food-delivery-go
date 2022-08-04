package restaurantusecase

import (
	"context"

	"github.com/daopmdean/food-delivery-go/common"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	Find(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	SoftDelete(ctx context.Context, id int) error
}

type deleteRestaurantUsecase struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantUsecase(store DeleteRestaurantStore) *deleteRestaurantUsecase {
	return &deleteRestaurantUsecase{
		store: store,
	}
}

func (usecase *deleteRestaurantUsecase) SoftDelete(ctx context.Context, id int) error {
	existedRst, err := usecase.store.Find(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotFindEntity(restaurantmodel.RestaurantEntityName, err)
	}

	if existedRst.Status == 0 {
		return common.ErrCannotFindEntity(restaurantmodel.RestaurantEntityName, err)
	}

	err = usecase.store.SoftDelete(ctx, id)
	if err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.RestaurantEntityName, err)
	}

	return nil
}

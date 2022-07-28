package restaurantusecase

import (
	"context"

	"github.com/daopmdean/food-delivery-go/common"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantUsecase struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantUsecase(store CreateRestaurantStore) *createRestaurantUsecase {
	return &createRestaurantUsecase{
		store: store,
	}
}

func (usecase *createRestaurantUsecase) CreateRestaurant(
	ctx context.Context,
	data *restaurantmodel.RestaurantCreate,
) error {
	err := data.Validate()
	if err != nil {
		return err
	}

	err = usecase.store.Create(ctx, data)
	if err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}

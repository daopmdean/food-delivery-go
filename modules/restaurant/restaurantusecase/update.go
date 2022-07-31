package restaurantusecase

import (
	"context"

	"github.com/daopmdean/food-delivery-go/common"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStorage interface {
	Update(
		ctx context.Context,
		id int,
		data *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantUsecase struct {
	store UpdateRestaurantStorage
}

func NewUpdateRestaurantUsecase(store UpdateRestaurantStorage) *updateRestaurantUsecase {
	return &updateRestaurantUsecase{
		store: store,
	}
}

func (usecase *updateRestaurantUsecase) UpdateRestaurant(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := usecase.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.RestaurantEntityName, err)
	}

	return nil
}

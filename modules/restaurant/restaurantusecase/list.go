package restaurantusecase

import (
	"context"

	"github.com/daopmdean/food-delivery-go/common"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
)

type ListRestaurantStore interface {
	List(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantUsecase struct {
	store ListRestaurantStore
}

func NewListRestaurantUsecase(store ListRestaurantStore) *listRestaurantUsecase {
	return &listRestaurantUsecase{
		store: store,
	}
}

func (usecase *listRestaurantUsecase) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := usecase.store.List(ctx, filter, paging)
	if err != nil {
		common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}

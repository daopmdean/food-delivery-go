package restaurantstorage

import (
	"context"

	"github.com/daopmdean/food-delivery-go/common"

	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Create(
	ctx context.Context,
	data *restaurantmodel.RestaurantCreate,
) error {
	db := s.db

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		common.ErrDb(err)
		return common.ErrDb(err)
	}

	return nil
}

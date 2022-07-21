package restaurantstorage

import (
	"context"
	"fmt"

	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Create(
	ctx context.Context,
	data *restaurantmodel.RestaurantCreate,
) error {
	db := s.db

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		return fmt.Errorf("error create restaurant: %s", err)
	}

	return nil
}

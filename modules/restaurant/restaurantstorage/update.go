package restaurantstorage

import (
	"context"
	"time"

	"github.com/daopmdean/food-delivery-go/common"

	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Update(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	db := s.db

	currentTime := time.Now()
	data.UpdatedAt = &currentTime

	if err := db.Table(data.TableName()).
		Where("id = ?", id).
		Updates(data).
		Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}

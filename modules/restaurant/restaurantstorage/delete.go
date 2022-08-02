package restaurantstorage

import (
	"context"

	"github.com/daopmdean/food-delivery-go/common"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
)

func (store *sqlStore) Delete(ctx context.Context, id int) error {
	db := store.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}

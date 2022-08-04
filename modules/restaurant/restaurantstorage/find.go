package restaurantstorage

import (
	"context"

	"github.com/daopmdean/food-delivery-go/common"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
)

func (store *sqlStore) Find(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	db := store.db
	var result restaurantmodel.Restaurant

	for _, item := range moreKeys {
		db = db.Preload(item)
	}

	if err := db.Where(condition).First(&result).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return &result, nil
}

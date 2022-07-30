package restaurantstorage

import (
	"context"
	"strconv"

	"github.com/daopmdean/food-delivery-go/common"
	"github.com/daopmdean/food-delivery-go/modules/restaurant/restaurantmodel"
)

func (store *sqlStore) List(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	db := store.db
	db = db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("status NOT IN (0)")

	if filter != nil {
		if filter.CategoryId != "" {
			cateId, err := strconv.Atoi(filter.CategoryId)
			if err == nil {
				db = db.Where("category_id = ?", cateId)
			}
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	db = db.Limit(paging.Limit)

	for _, k := range moreKeys {
		db = db.Preload(k)
	}

	if paging.Cursor > 0 {
		db = db.Where("id < ?", paging.Cursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Order("id DESC").Find(&result).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return result, nil
}

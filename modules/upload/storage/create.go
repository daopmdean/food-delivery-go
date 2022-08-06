package uploadstorage

import (
	"context"

	"github.com/daopmdean/food-delivery-go/common"
)

func (store *sqlStore) CreateImage(ctx context.Context, data *common.Image) error {
	db := store.db

	if err := db.Table(data.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}

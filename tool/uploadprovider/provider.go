package uploadprovider

import (
	"context"

	"github.com/daopmdean/food-delivery-go/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}

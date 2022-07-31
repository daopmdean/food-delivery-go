package restaurantmodel

import (
	"errors"
	"strings"
	"time"

	"github.com/daopmdean/food-delivery-go/common"
)

type RestaurantUpdate struct {
	Name             *string       `json:"name" gorm:"column:name;"`
	Address          *string       `json:"address" gorm:"column:addr;"`
	CityId           *int          `json:"-" gorm:"column:city_id;"`
	Lat              *float64      `json:"lat" gorm:"column:lat;"`
	Lng              *float64      `json:"lng" gorm:"column:lng;"`
	Logo             *common.Image `json:"logo" gorm:"column:logo;"`
	Cover            *common.Image `json:"cover" gorm:"column:cover;"`
	ShippingFeePerKm *float32      `json:"shipping_fee_per_km" gorm:"column:shipping_fee_per_km;"`
	UpdatedAt        *time.Time    `json:"-" gorm:"column:updated_at;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (r *RestaurantUpdate) Validate() error {
	if r.Name != nil {
		*r.Name = strings.TrimSpace(*r.Name)
	}

	if r.Address != nil {
		*r.Address = strings.TrimSpace(*r.Address)
	}

	if len(*r.Name) == 0 {
		return errors.New("restaurant name can not be empty")
	}

	if len(*r.Address) == 0 {
		return errors.New("restaurant address can not be empty")
	}

	return nil
}

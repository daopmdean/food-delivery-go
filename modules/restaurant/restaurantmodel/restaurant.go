package restaurantmodel

import "github.com/daopmdean/food-delivery-go/common"

const RestaurantEntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel  `json:",inline"`
	OwnerId          int           `json:"-" gorm:"column:owner_id;"`
	Name             string        `json:"name" gorm:"column:name;"`
	Address          string        `json:"address" gorm:"column:addr;"`
	CityId           int           `json:"-" gorm:"column:city_id;"`
	Lat              float64       `json:"lat" gorm:"column:lat;"`
	Lng              float64       `json:"lng" gorm:"column:lng;"`
	Logo             *common.Image `json:"logo" gorm:"column:logo;"`
	Cover            *common.Image `json:"cover" gorm:"column:cover;"`
	ShippingFeePerKm float32       `json:"shipping_fee_per_km" gorm:"column:shipping_fee_per_km;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

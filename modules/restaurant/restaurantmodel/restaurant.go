package restaurantmodel

import "github.com/daopmdean/food-delivery-go/common"

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	OwnerId         int     `json:"-" gorm:"column:owner_id;"`
	Name            string  `json:"name" gorm:"column:name;"`
	Address         string  `json:"address" gorm:"column:addr;"`
	CityId          int     `json:"-" gorm:"column:city_id;"`
	Lat             float64 `json:"lat"`
	Lng             float64 `json:"lng"`
	Logo            *common.Image
}

func (Restaurant) TableName() string {
	return "restaurants"
}

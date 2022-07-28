package restaurantmodel

import "github.com/daopmdean/food-delivery-go/common"

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	OwnerId         int     `json:"-"`
	Name            string  `json:"name"`
	Address         string  `json:"address"`
	CityId          int     `json:"-"`
	Lat             float64 `json:"lat"`
	Lng             float64 `json:"lng"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

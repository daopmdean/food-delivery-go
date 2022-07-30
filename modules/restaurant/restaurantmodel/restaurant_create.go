package restaurantmodel

import (
	"errors"
	"strings"
)

type RestaurantCreate struct {
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (r *RestaurantCreate) Validate() error {
	r.Name = strings.TrimSpace(r.Name)
	r.Address = strings.TrimSpace(r.Address)

	if len(r.Name) == 0 {
		return errors.New("restaurant name can not be empty")
	}

	if len(r.Address) == 0 {
		return errors.New("restaurant address can not be empty")
	}

	return nil
}

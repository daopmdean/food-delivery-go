package uploadmodel

import "github.com/daopmdean/food-delivery-go/common"

type Upload struct {
	common.SQLModel `json:",inline"`
	common.Image    `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

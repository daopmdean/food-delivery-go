package restaurantmodel

type Filter struct {
	CategoryId string `json:"category_id,omitempty" form:"category_id"`
}

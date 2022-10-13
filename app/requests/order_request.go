package requests

type OrderCreateRequest struct {
	Name        string `form:"name" json:"name" validate:"required"`
	Qty         int    `form:"qty" json:"qty" validate:"required,number"`
	Price       string `form:"price" json:"price" validate:"required"`
	Description string `form:"description" json:"description"`
}

type OrderUpdateRequest struct {
	ID          int    `json:"id" form:"id" validate:"required,omitempty,uuid"`
	Name        string `form:"name" json:"name" validate:"required"`
	Qty         int    `form:"qty" json:"qty" validate:"required,number"`
	Price       string `form:"price" json:"price" validate:"required"`
	Description string `form:"description" json:"description"`
}

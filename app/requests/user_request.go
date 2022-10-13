package requests

type UserCreateRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email" `
	Password string `json:"password" form:"password" validate:"required"`
}

type UserUpdateRequest struct {
	ID       int    `json:"id" form:"id" validate:"required,omitempty,uuid"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email" `
	Password string `json:"password" form:"password" validate:"required"`
}

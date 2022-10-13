package models

type User struct {
	BaseModel
	Name     string  `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Email    string  `gorm:"type:varchar(255);unique" json:"email" validate:"required,email"`
	Password string  `gorm:"type:varchar(255)" json:"password" validate:"required"`
	Orders   []Order `json:"orders"`
}

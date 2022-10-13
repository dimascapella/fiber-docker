package models

type Order struct {
	BaseModel
	Name        string `gorm:"type:varchar(255); unique" json:"name" validate:"required"`
	Qty         int    `gorm:"not null" json:"qty" validate:"required,number"`
	Price       string `gorm:"type:varchar(255)" json:"price" validate:"required"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	UserID      int    `gorm:"index" json:"user_id"`
	User        User   `gorm:"foreignKey:user_id;references:id"`
}

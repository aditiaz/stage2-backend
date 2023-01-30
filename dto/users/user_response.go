package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Status   string `gorm:"type: varchar(255)" json:"status" validate:"required"`
	Gender   string `gorm:"type: varchar(255)" json:"gender" validate:"required"`
	Address  string `gorm:"type: varchar(255)" json:"address" validate:"required"`
}

package usersdto

type CreateUserRequest struct {
	Name     string `gorm:"type: varchar(255)" json:"name" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Status   string `gorm:"type: varchar(255)" json:"status" validate:"required"`
	Gender   string `gorm:"type: varchar(255)" json:"gender" validate:"required"`
	Address  string `gorm:"type: varchar(255)" json:"address" validate:"required"`
}
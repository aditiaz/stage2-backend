package transactiondto

type Request_Transaction struct {
	Check_In       string `gorm:"type: varchar(255)" json:"check_in" form:"check_in"`
	Check_Out      string `gorm:"type: varchar(255)" json:"check_out" form:"check_out"`
	House_Id       int    `json:"house_id" form:"house_id"`
	User_Id        int    `json:"user_id" form:"user_id"`
	Total          int    `json:"total" form:"total"`
	Status_Payment string `json:"status_payment" form:"status_payment"`
	Image_Payment  string `json:"image_payment" form:"image_payment"`
}

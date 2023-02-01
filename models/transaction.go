package models

type Transaction struct {
	ID             int    `json:"id"`
	Check_In       string `json:"check_in"`
	Check_Out      string `json:"check_out"`
	House_Id       int    `json:"house_id"`
	House          House
	User_Id        int `json:"user_id"`
	User           User
	Total          int    `json:"total"`
	Status_Payment string `json:"status_payment"`
	Image_Payment  string `json:"image_payment"`
}

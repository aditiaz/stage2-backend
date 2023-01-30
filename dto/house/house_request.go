package housedto

type Create_Request_Property struct {
	Name_House    string `json:"name_property" form:"name_property" validate:"required"`
	City          string `json:"city" form:"city" validate:"required"`
	Address_House string `json:"address_property" form:"address_property" validate:"required"`
	Price         int    `json:"price" form:"price" validate:"required"`
	Type_Of_Rent  string `json:"type_of_rent" form:"type_of_rent" validate:"required"`
	Amenities     string `json:"amenities" form:"amenities" validate:"required"`
	Bed_Room      int    `json:"bed_room" form:"bed_room" validate:"required"`
	Bath_Room     int    `json:"bath_room" form:"bath_room" validate:"required"`
	Description   string `json:"description" form:"description" validate:"required"`
	Image_House   string `json:"image_property" form:"image_property"`
}
type Update_Request_Property struct {
	Name_House    string `json:"name_property" form:"name_property" validate:"required"`
	City          string `json:"city" form:"city" validate:"required"`
	Address_House string `json:"address_property" form:"address_property" validate:"required"`
	Price         int    `json:"price" form:"price" validate:"required"`
	Type_Of_Rent  string `json:"type_of_rent" form:"type_of_rent" validate:"required"`
	Amenities     string `json:"amenities" form:"amenities" validate:"required"`
	Bed_Room      int    `json:"bed_room" form:"bed_room" validate:"required"`
	Bath_Room     int    `json:"bath_room" form:"bath_room" validate:"required"`
	Description   string `json:"description" form:"description" validate:"required"`
	Image_House   string `json:"image_house" form:"image_house" `
}
package models

type House struct {
	ID            int    `json:"id"`
	Name_House    string `json:"name_property"`
	City          string `json:"city"`
	Address_House string `json:"address_property"`
	Price         int    `json:"price"`
	Type_Of_Rent  string `json:"type_of_rent"`
	Amenities     string `json:"amenities"`
	Bed_Room      int    `json:"bed_room"`
	Bath_Room     int    `json:"bath_room"`
	Description   string `json:"description"`
	Image_House   string `json:"image_property"`
}

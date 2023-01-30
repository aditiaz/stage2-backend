package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	housedto "task1/dto/house"
	dto "task1/dto/result"
	"task1/models"
	"task1/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerHouse struct {
	HouseRepository repositories.HouseRepository
}

func HandlerProperty(HouseRepository repositories.HouseRepository) *handlerHouse {
	return &handlerHouse{HouseRepository}
}

func (h *handlerHouse) FindHouses(w http.ResponseWriter ,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	houses, err := h.HouseRepository.FindHouses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK,Data:houses}
	json.NewEncoder(w).Encode(response)
	fmt.Println(houses)
}

func (h *handlerHouse)GetHouse(w http.ResponseWriter ,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id,_ := strconv.Atoi(mux.Vars(r)["id"])

	house,err := h.HouseRepository.GetHouse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// to show succeded data
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK,Data: house}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHouse)CreateHouse(w http.ResponseWriter ,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// requesting
	price,_ := strconv.Atoi(r.FormValue("price"))
	bed_room,_ := strconv.Atoi(r.FormValue("bed_room"))
	bath_room,_ := strconv.Atoi(r.FormValue("bath_room"))

	// fmt.Println(r.FormValue("image"), "Halo")
	request := housedto.Create_Request_Property {
		Name_House : r.FormValue("name_property"),
		City  : r.FormValue("city"),
	    Address_House : r.FormValue("address_property"),
		Price : price,
		Type_Of_Rent : r.FormValue("type_of_rent"),
		Amenities : r.FormValue("amenities"),
		Bed_Room : bed_room,
        Bath_Room : bath_room,
		Description  : r.FormValue("description"),
		// Image_House : r.FormValue("image_property"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// image
	// dataContext := r.Context().Value("Error")
	// var filename string
	// if dataContext == nil {
	// 	dataContext := r.Context().Value("dataFile")
	// 	filename = dataContext.(string)
	// }

	house := models.House {
		Price : request.Price,
        Bed_Room : request.Bed_Room,
		Bath_Room   : request.Bath_Room,
		Name_House  : request.Name_House,
		City   : request.City,
		Address_House : request.Address_House,
		Type_Of_Rent : request.Type_Of_Rent,
		Amenities  : request.Amenities,
		Description : request.Description,
		// Image_House: filename,
	}


	data, err := h.HouseRepository.CreateHouse(house)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	// get data
	houseGet, err := h.HouseRepository.GetHouse(data.ID)
	// houseGet.Image_House = os.Getenv("PATH_FILE") + houseGet.Image_House
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
	// success
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertPropertyResponse(houseGet)}
	json.NewEncoder(w).Encode(response)
  }
  
  // Update data
  func (h *handlerHouse) UpdateHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	// params
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// get data
	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
	// validation
	if r.FormValue("name_house") != "" {
	  house.Name_House = r.FormValue("name_house")
	}
	Price, _ := strconv.Atoi(r.FormValue("price"))
	if Price != 0 {
	  house.Price = Price
	}
	if r.FormValue("city") != "" {
	  house.City = r.FormValue("city")
	}
	if r.FormValue("address_house") != "" {
	  house.Address_House = r.FormValue("address_house")
	}
	if r.FormValue("type_of_rent") != "" {
	  house.Type_Of_Rent = r.FormValue("type_of_rent")
	}
	bed_room, _ := strconv.Atoi(r.FormValue("bed_room"))
	if bed_room != 0 {
	  house.Bed_Room = bed_room
	}
	bath_room, _ := strconv.Atoi(r.FormValue("bath_room"))
	if bath_room != 0 {
	  house.Bath_Room = bath_room
	}

	if r.FormValue("Amenities") != "" {
	  house.Amenities = r.FormValue("Amenities")
	}
	if r.FormValue("Description") != "" {
	  house.Description = r.FormValue("Description")
	}
  
	// image
	dataContex := r.Context().Value("Error")
  
	// fmt.Println(dataContex)
	if dataContex == nil {
	  // image
	  dataContex := r.Context().Value("dataFile")
	  filename := dataContex.(string)
	  house.Image_House = filename
	}
  
	// update data
	data, err := h.HouseRepository.UpdateHouse(house)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	// get data
	// houseInserted,err :=
	houseInserted,err := h.HouseRepository.GetHouse(data.ID)
	houseInserted.Image_House = os.Getenv("PATH_FILE") + houseInserted.Image_House
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
	// success
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertPropertyResponse(houseInserted)}
	json.NewEncoder(w).Encode(response)
}
  
  

  
  func convertPropertyResponse(r models.House) housedto.Response_Property {
	return housedto.Response_Property{
	  ID:               r.ID,
	  Price:            r.Price,
	  Bed_Room:         r.Bed_Room,
	  Bath_Room:        r.Bath_Room,
	  Name_House:    r.Name_House,
	  City:             r.City,
	  Address_House: r.Address_House,
	  Type_Of_Rent:     r.Type_Of_Rent,
	  Amenities:        r.Amenities,
	  Image_House:   r.Image_House,
	  Description:      r.Description,
	}
  }

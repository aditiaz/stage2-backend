package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	// housedto "task1/dto/house"
	dto "task1/dto/result"
	transactiondto "task1/dto/transaction"
	"task1/models"
	"task1/repositories"

	// "github.com/go-playground/locales/id"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransactionRepository(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransactions(w http.ResponseWriter ,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	houses, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK,Data:houses}
	json.NewEncoder(w).Encode(response)
	fmt.Println(houses)
}

func (h *handlerTransaction)GetTransaction(w http.ResponseWriter ,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id,_ := strconv.Atoi(mux.Vars(r)["id"])

	house,err := h.TransactionRepository.GetTransaction(id)
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

func (h *handlerTransaction)CreateTransaction(w http.ResponseWriter ,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Get dataFile from midleware and store to filename variable here ...
	dataContext := r.Context().Value("dataFile") 
	filename := dataContext.(string)            

	// requesting
	House_Id,_ := strconv.Atoi(r.FormValue("house_id"))
	User_Id,_ := strconv.Atoi(r.FormValue("user_id"))
	Total,_ := strconv.Atoi(r.FormValue("total"))
	// fmt.Println(r.FormValue("image"), "Halo")
	request := transactiondto.Request_Transaction {
		Check_In: r.FormValue("check_in"),
		Check_Out: r.FormValue("check_out"),
		House_Id: House_Id,
		User_Id: User_Id,
		Total: Total,
		Status_Payment: r.FormValue("status_payment"),
		Image_Payment: r.FormValue("image_payment"),	
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	
	transaction := models.Transaction {
    Check_In: request.Check_In,
	Check_Out: request.Check_Out,
	House_Id: request.House_Id,
	User_Id:userId  ,
	Total: request.Total,
	Status_Payment: request.Status_Payment,
	Image_Payment: filename,
	}


	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	// get data
	transactionGet, err := h.TransactionRepository.GetTransaction(data.ID)
	// houseGet.Image_House = os.Getenv("PATH_FILE") + houseGet.Image_House
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
	// success
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertTransactionResponse(transactionGet)}
	json.NewEncoder(w).Encode(response)
  }
  
  
 
 
 
 
  // Update data
  func (h *handlerTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

  
	// params
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// get data
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
	// validation
	if r.FormValue("check_in") != "" {
		transaction.Check_In = r.FormValue("check_in")
	}
	
	if r.FormValue("check_out") != "" {
	  transaction.Check_Out = r.FormValue("check_out")
	}
	House_id, _ := strconv.Atoi(r.FormValue("house_id"))
	if House_id != 0 {
	  transaction.House_Id =House_id
	}
	User_id, _ := strconv.Atoi(r.FormValue("user_id"))
	if User_id != 0 {
	  transaction.User_Id = User_id
	}
	Total, _ := strconv.Atoi(r.FormValue("total"))
	if Total != 0 {
	  transaction.Total = Total
	}
	if r.FormValue("status_payment") != "" {
		transaction.Status_Payment = r.FormValue("status_payment")
	  }
	
	
	dataContex := r.Context().Value("Error")  
	// fmt.Println(dataContex)
	if dataContex == nil {
	  dataContex := r.Context().Value("dataFile")
	  filename := dataContex.(string)
	  transaction.Image_Payment = filename
	}
  
	// update data
	data, err := h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	// get data
	// houseInserted,err :=
	transactionInserted,err := h.TransactionRepository.GetTransaction(data.ID)
	transactionInserted.Image_Payment = os.Getenv("PATH_FILE") + transactionInserted.Image_Payment
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
	// success
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertTransactionResponse(transactionInserted)}
	json.NewEncoder(w).Encode(response)
}
  
// func (h *handlerTransaction) DeleteHouse(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	house, err := h.HouseRepository.GetHouse(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	data, err := h.HouseRepository.DeleteHouse(house)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: convertPropertyResponse(data)}
// 	json.NewEncoder(w).Encode(response)
// }
  

  
  func convertTransactionResponse(r models.Transaction)transactiondto.Response_Transaction{
	  
	
	  return transactiondto.Response_Transaction{
		ID: r.ID,
		Check_In      : r.Check_In,
		Check_Out      : r.Check_Out,
		House_Id     : r.House_Id,
		User_Id      : r.User_Id,
		Total        : r.Total,
		Status_Payment : r.Status_Payment,
		Image_Payment  : r.Image_Payment,
	}
  }

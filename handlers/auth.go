package handlers

// Dont forget import required packages this below ...
import (
	"encoding/json"
	"net/http"
	authdto "task1/dto/auth"
	dto "task1/dto/result"
	"task1/models"
	"task1/pkg/bcrypt"
	jwtToken "task1/pkg/jwt"

	"github.com/golang-jwt/jwt/v4"

	// "encoding/json"
	"fmt"
	"log"

	// "net/http"
	"task1/repositories"
	"time"

	"github.com/go-playground/validator/v10"
)

// handlerAuth struct here ...
// ngambil fungsi2 query dari folder repositories auth
type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

// HandlerAuth function here ...
func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	user := models.User{
		Fullname:     request.Name,
		Email:    request.Email,
		Password: password,
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseAuth(data)}
	json.NewEncoder(w).Encode(response)
	w.Header().Set("Content-Type", "application/json")
}

// Create Login method here ...
// Login adalah fungsi query dari folder repositories
func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

// mengambil struct Login request :
	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
// memasukan steuct Login request ke dalam struct models.User
	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	// Check email
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	loginResponse := authdto.LoginResponse{
		Name:     user.Fullname,
		Email:    user.Email,
		Password: user.Password,
		Token:    token,
	}

	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)

}

func convertResponseAuth(u models.User) authdto.RegisterResponse {
	return authdto.RegisterResponse{
		ID:       u.ID,
		Name:     u.Fullname,
		Email:    u.Email,
		Password: u.Password,
		// Status: u.Status,
		// Gender: u.Gender,
		// Address: u.Address,
	
	}
}

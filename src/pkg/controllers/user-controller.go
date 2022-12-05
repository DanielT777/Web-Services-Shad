package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DanielT777/Web-Services-Shad/src/pkg/models"
	"github.com/DanielT777/Web-Services-Shad/src/pkg/utils"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var newUser models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	userDetails := models.GetUserByID(id)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(CreateUser.Password), 8)
	if err != nil {
		fmt.Println(err)
	}
	CreateUser.Password = string(hashedPassword)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Target object after decoding:", CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	utils.ParseBody(r, &newUser)
	// check for non-empty fields
	userDetails := models.GetUserByID(id)
	if newUser.FirstName != "" {
		userDetails.FirstName = newUser.FirstName
	}
	if newUser.LastName != "" {
		userDetails.LastName = newUser.LastName
	}
	if newUser.PseudoName != "" {
		userDetails.PseudoName = newUser.PseudoName
	}
	if newUser.Email != "" {
		userDetails.Email = newUser.Email
	}
	newUser.ID = uint(id)
	newUser.UpdateUser()
	res, _ := json.Marshal(newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	models.DeleteUser(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

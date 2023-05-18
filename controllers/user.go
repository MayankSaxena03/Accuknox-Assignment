package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MayamkSaxena03/Accuknox/models"
	"github.com/MayamkSaxena03/Accuknox/utility"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.UserData
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || !user.ValidateSignupBody() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}

	exists, err := models.CheckUserExists(r.Context(), user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if exists {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode("User already exists")
		return
	}

	userCollection := models.GetUserCollection()
	_, err = userCollection.InsertOne(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("User created successfully")
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.UserData
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || !user.ValidateLoginBody() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}

	exists, err := models.CheckUserExists(r.Context(), user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User does not exist")
		return
	}

	userCollection := models.GetUserCollection()
	var dbUser models.UserData
	err = userCollection.FindOne(r.Context(), models.UserData{Email: user.Email}).Decode(&dbUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if dbUser.Password != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid password")
		return
	}

	token, err := utility.GenerateTokenForUser(dbUser.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(
		bson.M{
			"sid": token,
		},
	)
}

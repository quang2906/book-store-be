package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/quang2906/book_store_be/database"
	jwt "github.com/quang2906/book_store_be/jwtoftris"
	"github.com/quang2906/book_store_be/model"
	"github.com/quang2906/book_store_be/util"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	database.Connect()
	defer database.DB.Clauses()

	var authDetails model.Authentication

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err util.Error
		err = util.SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var authUser model.User
	database.DB.Where("email = 	?", authDetails.Email).First(&authUser)

	if authUser.Email == "" {
		var err util.Error
		err = util.SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	check := model.CheckPasswordHash(authDetails.Password, string(authUser.Password))

	if !check {
		var err util.Error
		err = util.SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := jwt.GenerateJWT(authUser.Email, authUser.Role)

	http.SetCookie(w, &http.Cookie{
		Name:     "Token",
		Value:    validToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	})

	if err != nil {
		var err util.Error
		err = util.SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	var token model.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "logout successfully",
	})
}

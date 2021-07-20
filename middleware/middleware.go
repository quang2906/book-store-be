package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/quang2906/book_store_be/util"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			var err util.Error
			err = util.SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte("trisdepzai")

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err util.Error
			err = util.SetError(err, "Your Token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return

			}
		}
		var reserr util.Error
		reserr = util.SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}

// func IsAuthorized2(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := r.Cookie("token")
// 		if err != nil {
// 			if err == http.ErrNoCookie {
// 				http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 				return
// 			}
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		id, _ := jwt.ParseJwt(cookie.Value)
// 		userId, _ := strconv.Atoi(id)
// 		user := model.User{
// 			ID: int64(userId),
// 		}

// 		database.Connect()
// 		defer database.DB.Clauses()

// 		row, _ := database.DB.Raw("select r.name from users u "
// 			"join roles r on r.id = u.role_id where u.id = ?", user.ID)
// 		if row.Next() {
// 			row.Scan(&user.Role)
// 		}
// 		fmt.Println("role: ", user.Role)

// 		if user.Role != "ADMIN" {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			json.NewEncoder(w).Encode(map[string]string{
// 				"message": "access denied",
// 			})
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

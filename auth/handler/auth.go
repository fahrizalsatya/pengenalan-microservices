package handler

import (
	"net/http"

	"github.com/FahrizalSatya/pengenalan-microservices/utils"
)

//ValidateAuth for validating user connect to database
func ValidateAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
		return
	}

	if authToken != "asdfghjk" {
		utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
		return
	}
	utils.WrapAPISuccess(w, r, "success", 200)
}

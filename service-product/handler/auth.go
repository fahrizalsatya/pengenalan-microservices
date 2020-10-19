package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/FahrizalSatya/pengenalan-microservices/service-product/config"
	"github.com/FahrizalSatya/pengenalan-microservices/utils"
)

//AuthMiddleware represents AuthService
type AuthMiddleware struct {
	AuthService config.AuthService
}

//ValidateAuth from auth Middleware
func (auth *AuthMiddleware) ValidateAuth(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := http.NewRequest("POST", auth.AuthService.Host+"/admin-auth", nil)
		if err != nil {
			utils.WrapAPIError(w, r, "failed to create request : "+err.Error(), http.StatusInternalServerError)
			return
		}

		request.Header = r.Header
		authResponse, err := http.DefaultClient.Do(request)
		if err != nil {
			utils.WrapAPIError(w, r, "validate auth failed : "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer authResponse.Body.Close()
		body, err := ioutil.ReadAll(authResponse.Body)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		var authResult map[string]interface{}
		err = json.Unmarshal(body, &authResult)

		if authResponse.StatusCode != 200 {
			utils.WrapAPIError(w, r, authResult["error_details"].(string), authResponse.StatusCode)
			return
		}

		nextHandler(w, r)
	}
}

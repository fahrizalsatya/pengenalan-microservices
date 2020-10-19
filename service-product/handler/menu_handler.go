package handler

import (
	"net/http"

	"github.com/FahrizalSatya/pengenalan-microservices/utils"
)

//AddMenuHandler handle add menu
func AddMenuHandler(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", 200)
}

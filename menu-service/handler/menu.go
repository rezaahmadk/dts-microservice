package handler

import (
	"net/http"

	"github.com/rezaahmadk/dts-microservice/utils"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
}

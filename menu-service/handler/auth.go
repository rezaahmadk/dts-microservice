package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/context"
	"github.com/rezaahmadk/dts-microservice/menu-service/config"
	"github.com/rezaahmadk/dts-microservice/menu-service/entity"
	"github.com/rezaahmadk/dts-microservice/utils"
)

type AuthMiddleware struct {
	AuthService config.AuthService
}

func (auth *AuthMiddleware) ValidateAuth(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := http.NewRequest("POST", auth.AuthService.Host+"/auth/validate", nil)
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

		var authResult entity.AuthResponse
		err = json.Unmarshal(body, &authResult)

		if authResponse.StatusCode != 200 {
			utils.WrapAPIError(w, r, authResult.ErrorDetails, authResponse.StatusCode)
			return
		}

		context.Set(r, "user", authResult.Data.Username)

		nextHandler(w, r)
	}
}

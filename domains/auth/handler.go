package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/layerpro/easy-print-hub-backend/utils"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var data SignIn
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf(`Error get body %v`, err)
		utils.ResponseError(w, http.StatusBadRequest, utils.ErrBodyDecode)
		return
	}
	message, err := utils.Validator(data)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, message)
		return
	}

	signIn, err := h.service.SignIn(data)
	if err == ErrWrongEmailOrPassword {
		utils.ResponseError(w, http.StatusBadRequest, utils.ErrWrongEmailOrPassword)
		return
	}
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, utils.ErrInternalServerError)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, utils.SuccessOk, signIn)
}

func (h Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	accessToken := utils.GetBearerToken(r)
	if accessToken == "" {
		utils.ResponseError(w, http.StatusBadRequest, utils.ErrGetProfile)
		return
	}

	err := h.service.SignOut(accessToken)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, utils.ErrUnauthorized)
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, utils.SuccessOk, nil)
}

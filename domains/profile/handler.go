package profile

import (
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

func (h Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	user, err := utils.UserFromContext(r.Context())
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, utils.ErrGetProfile)
		return
	}

	profile, err := h.service.GetProfile(user.ID)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, utils.ErrInternalServerError)
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, utils.SuccessOk, profile)
}

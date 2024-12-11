package handler

import (
	"net/http"
	"strconv"

	"fanchann/elegant_err/services"
	"fanchann/elegant_err/utils"
)

type UsersHandler struct {
	userService *services.UsersServices
}

func NewUsersHandlerImpl(userService *services.UsersServices) *UsersHandler {
	return &UsersHandler{userService}
}

func (uh *UsersHandler) SearchUserById(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")
	userIdInt := utils.CatchError(strconv.Atoi(userId))
	if userIdInt == 0 {
		return
	}
	user, err := uh.userService.SearchUserById(userIdInt)

	utils.ErrorAPIHandler(user, err, w)
}

package api

import (
	"net/http"
	"strconv"

	m "github.com/KasissLiu/go-webserver/models"
)

type uController struct{}

var User userController = userController{}

func (u *userController) GetUserById(w http.ResponseWriter, h *http.Request) {
	h.ParseForm()

	id := h.Form.Get("id")
	if id == "" {
		responseJson(1, "id未获取到", "", w)
		return
	}
	userId, _ := strconv.Atoi(id)
	user, err := m.UserModel.GetUserById(userId)
	if err != nil {
		responseJson(1, "未搜索到用户", "", w)
		return
	}
	responseJson(0, "success", user, w)
}

func (u *userController) GetUserAll(w http.ResponseWriter, h *http.Request) {
	users := m.UserModel.GetUserAll()
	responseJson(0, "success", users, w)
}

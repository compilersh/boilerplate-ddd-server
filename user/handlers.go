package user

import (
	"fmt"
	"net/http"

	"github.com/compilersh/boilerplate-ddd-server/reqres"
	"github.com/go-chi/chi"
)

type Handler struct {
	service *UserService
}

func NewHandler(service *UserService) *Handler {
	return &Handler{
		service: service,
	}
}

type UserReq struct {
	Username string `json:"username"`
}

type userRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// Validate implements the Validator interface from reqres package.
func (u UserReq) Validate() error {
	if u.Username == "" {
		return fmt.Errorf("username cannot be empty")
	}
	return nil
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req UserReq
	if err := reqres.DecodeJSON(r, &req); err != nil {
		reqres.ResJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		reqres.ResJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	reqres.ResJSON(w, http.StatusCreated, user.ToUserRes())
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.service.GetUser(id)
	if err != nil {
		reqres.ResJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	reqres.ResJSON(w, http.StatusOK, user.ToUserRes())
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		reqres.ResJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	usersRes := make([]userRes, 0, len(users))
	for _, u := range users {
		usersRes = append(usersRes, u.ToUserRes())
	}

	reqres.ResJSON(w, http.StatusOK, usersRes)
}

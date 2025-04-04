package netInterface

import (
	"encoding/json"
	"net/http"

	"PhBook/domen"
	"PhBook/server/jwt"
	"PhBook/userCase"
)

type AuthHandlers struct {
	pb *userCase.PhoneBook
}

func NewAuthHandlers(pb *userCase.PhoneBook) *AuthHandlers {
	return &AuthHandlers{
		pb: pb,
	}
}

func (h *AuthHandlers) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req domen.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		http.Error(w, domen.ErrOperationFailed.Error(), http.StatusBadRequest)
		return
	}

	if req.Username == "" {

		http.Error(w, domen.ErrEmptyUsername.Error(), http.StatusBadRequest)
		return
	}

	if req.Password == "" {

		http.Error(w, domen.ErrEmptyPassword.Error(), http.StatusBadRequest)
		return
	}

	if err := h.pb.RegisterUser(req.Username, req.Password); err != nil {

		switch err {
		case domen.ErrUserExists:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func (h *AuthHandlers) HandleAuth(w http.ResponseWriter, r *http.Request) {
	var req domen.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		http.Error(w, domen.ErrOperationFailed.Error(), http.StatusBadRequest)
		return
	}

	userID, err := h.pb.AuthUser(req.Username, req.Password)
	if err != nil {

		switch err {
		case domen.ErrInvalidCredentials, domen.ErrUserNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	token, err := jwt.GenerateJWT(userID)
	if err != nil {

		http.Error(w, domen.ErrOperationFailed.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":  token,
		"userID": userID,
	})
}

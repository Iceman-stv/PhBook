package user_http

import (
    "encoding/json"
    "net/http"

    "PhBook/userCase"
    "PhBook/server/utils"
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
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    	
        http.Error(w, "Неправильный запрос (регистрация)", http.StatusBadRequest)
        return
    }

    if err := h.pb.RegisterUser(req.Username, req.Password); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Пользователь зарегистрирован"))
}

func (h *AuthHandlers) HandleAuth(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    	
        http.Error(w, "Неправильный запрос (Аутентификация)", http.StatusBadRequest)
        return
    }

    userID, err := h.pb.AuthUser(req.Username, req.Password)
    if err != nil {
    	
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateJWT(userID)
    if err != nil {
    	
        http.Error(w, "Ошибка при генерации токена", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
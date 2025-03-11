package user_http

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "PhBook/userCase"
)

type ContactHandlers struct {
    pb *userCase.PhoneBook
}

func NewContactHandlers(pb *userCase.PhoneBook) *ContactHandlers {
    return &ContactHandlers{
        pb: pb,
    }
}

func (h *ContactHandlers) HandleAddContact(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Name  string `json:"name"`
        Phone string `json:"phone"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    	
        http.Error(w, "Неправильный запрос (Добавление контакта)", http.StatusBadRequest)
        return
    }

    userID := r.Context().Value("userID").(int)
    if err := h.pb.AddContact(userID, req.Name, req.Phone); err != nil {
    	
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Контакт добавлен"))
}

func (h *ContactHandlers) HandleGetContacts(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(int)
    contacts, err := h.pb.GetContacts(userID)
    if err != nil {
    	
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(contacts)
}

func (h *ContactHandlers) HandleDeleteContact(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(int)
    name := mux.Vars(r)["name"]
    if err := h.pb.DelContact(userID, name); err != nil {
    	
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Контакт удалён"))
}

func (h *ContactHandlers) HandleFindContact(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(int)
    name := r.URL.Query().Get("name")
    contacts, err := h.pb.FindContact(userID, name)
    if err != nil {
    	
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(contacts)
}
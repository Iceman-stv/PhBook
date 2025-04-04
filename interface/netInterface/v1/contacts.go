package netInterface

import (
	"encoding/json"
	"net/http"

	"PhBook/domen"
	"PhBook/userCase"

	"github.com/gorilla/mux"
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
	w.Header().Set("Content-Type", "application/json")

	var contact domen.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {

		http.Error(w, domen.ErrOpertionFailed.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(int)
	contact.UserID = userID

	if err := h.pb.AddContact(contact.UserID, contact.Name, contact.Phone); err != nil {

		switch err {
		case domen.ErrContactExists:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contact)
}

func (h *ContactHandlers) HandleGetContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value("userID").(int)
	if !ok {

		http.Error(w, domen.ErrOpertionFailed.Error(), http.StatusBadRequest)
		return
	}

	contacts, err := h.pb.GetContacts(userID)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contacts)
}

func (h *ContactHandlers) HandleDeleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value("userID").(int)
	if !ok {

		http.Error(w, domen.ErrOpertionFailed.Error(), http.StatusBadRequest)
		return
	}

	name := mux.Vars(r)["name"]
	if err := h.pb.DelContact(userID, name); err != nil {

		switch err {
		case domen.ErrContactNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Contact deleted successfully"})
}

func (h *ContactHandlers) HandleFindContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value("userID").(int)
	if !ok {

		http.Error(w, domen.ErrOpertionFailed.Error(), http.StatusBadRequest)
		return
	}

	name := r.URL.Query().Get("name")
	contacts, err := h.pb.FindContact(userID, name)
	if err != nil {

		switch err {
		case domen.ErrContactNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contacts)
}

package community

import (
	"encoding/json"
	"net/http"
)

type CreateCommunityPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CommunityHandler struct {
	service *CommunityService
}

func NewCommunityHandler(service *CommunityService) *CommunityHandler {
	return &CommunityHandler{service: service}
}

func (h *CommunityHandler) CreateCommunityHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var payload CreateCommunityPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	community, err := h.service.CreateCommunity(payload.Name, payload.Description, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(community)
}

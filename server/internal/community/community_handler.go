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

// GetCommunityHandler menangani permintaan untuk detail komunitas
func (h *CommunityHandler) GetCommunityHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	communityDetails, err := h.service.GetCommunityDetails(name)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			http.Error(w, "Community not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(communityDetails)
}

// JoinCommunityHandler menangani permintaan untuk bergabung ke komunitas
func (h *CommunityHandler) JoinCommunityHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	communityName := r.PathValue("name")

	// Kita perlu mendapatkan ID komunitas dari namanya
	community, err := h.service.GetCommunityDetails(communityName) // Kita bisa buat service yang lebih ringan nanti
	if err != nil {
		http.Error(w, "Community not found", http.StatusNotFound)
		return
	}

	if err := h.service.JoinCommunity(userID, community.ID); err != nil {
		http.Error(w, "Failed to join community", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Successfully joined community"})
}

// LeaveCommunityHandler menangani permintaan untuk meninggalkan komunitas
func (h *CommunityHandler) LeaveCommunityHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	communityName := r.PathValue("name")

	community, err := h.service.GetCommunityDetails(communityName)
	if err != nil {
		http.Error(w, "Community not found", http.StatusNotFound)
		return
	}

	if err := h.service.LeaveCommunity(userID, community.ID); err != nil {
		http.Error(w, "Failed to leave community", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Successfully left community"})
}

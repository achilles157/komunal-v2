package community

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CreateCommunityPayload struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"` // Tambahkan slug
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

	community, err := h.service.CreateCommunity(payload.Name, payload.Slug, payload.Description, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(community)
}

// GetCommunityHandler
func (h *CommunityHandler) GetCommunityHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)  // Gunakan mux.Vars
	name := vars["name"] // Ambil 'name' dari vars

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
	vars := mux.Vars(r) // Gunakan mux.Vars
	communityName := vars["name"]

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
	vars := mux.Vars(r)
	communityName := vars["name"]

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

// GetUserCommunitiesHandler menangani permintaan untuk mengambil komunitas milik user
func (h *CommunityHandler) GetUserCommunitiesHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	communities, err := h.service.GetUserCommunities(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve user communities", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(communities)
}

// DeleteCommunityHandler menangani permintaan untuk menghapus komunitas
func (h *CommunityHandler) DeleteCommunityHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r) // Gunakan mux.Vars
	communityName := vars["name"]

	// Dapatkan detail komunitas untuk mendapatkan ID-nya
	community, err := h.service.GetCommunityDetails(communityName)
	if err != nil {
		http.Error(w, "Community not found", http.StatusNotFound)
		return
	}

	// Service akan memvalidasi apakah userID adalah kreator
	if err := h.service.DeleteCommunity(community.ID, userID); err != nil {
		// Jika service mengembalikan error, kemungkinan user bukan kreator
		http.Error(w, err.Error(), http.StatusForbidden) // 403 Forbidden
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Community successfully deleted"})
}

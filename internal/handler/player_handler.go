package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gofootball/internal/model"
)

// GetAllPlayers handles GET /players
func (h *Handler) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := h.PlayerRepo.FindAll()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, players)
}

// CreatePlayer handles POST /player
func (h *Handler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	player := model.Player{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&player); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := h.PlayerRepo.Create(&player); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, player)
}

// GetPlayer handles GET /player/{name}
func (h *Handler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	player, err := h.PlayerRepo.FindByName(name)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, player)
}

// UpdatePlayer handles PUT /player/{name}
func (h *Handler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	player, err := h.PlayerRepo.FindByName(name)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(player); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := h.PlayerRepo.Update(player); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, player)
}

// DeletePlayer handles DELETE /player/{name}
func (h *Handler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	player, err := h.PlayerRepo.FindByName(name)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	if err := h.PlayerRepo.Delete(player); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// DisablePlayer handles PUT /player/{name}/disable
func (h *Handler) DisablePlayer(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	player, err := h.PlayerRepo.FindByName(name)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	player.Disable()
	if err := h.PlayerRepo.Update(player); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, player)
}

// EnablePlayer handles PUT /player/{name}/enable
func (h *Handler) EnablePlayer(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	player, err := h.PlayerRepo.FindByName(name)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	player.Enable()
	if err := h.PlayerRepo.Update(player); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, player)
}

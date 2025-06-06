package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gofootball/internal/model"
)

// GetAllClubs handles GET /clubs
func (h *Handler) GetAllClubs(w http.ResponseWriter, r *http.Request) {
	clubs, err := h.ClubRepo.FindAll()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, clubs)
}

// CreateClub handles POST /club
func (h *Handler) CreateClub(w http.ResponseWriter, r *http.Request) {
	club := model.Club{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&club); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := h.ClubRepo.Create(&club); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, club)
}

// GetClub handles GET /club/{name}
func (h *Handler) GetClub(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	club, err := h.ClubRepo.FindByName(name)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, club)
}

// UpdateClub handles PUT /club/{name}
func (h *Handler) UpdateClub(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	club, err := h.ClubRepo.FindByName(name)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(club); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := h.ClubRepo.Update(club); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, club)
}

// DeleteClub handles DELETE /club/{name}
func (h *Handler) DeleteClub(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	club, err := h.ClubRepo.FindByName(name)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	if err := h.ClubRepo.Delete(club); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// DisableClub handles PUT /club/{name}/disable
func (h *Handler) DisableClub(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	club, err := h.ClubRepo.FindByName(name)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	club.Disable()
	if err := h.ClubRepo.Update(club); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, club)
}

// EnableClub handles PUT /club/{name}/enable
func (h *Handler) EnableClub(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	club, err := h.ClubRepo.FindByName(name)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	club.Enable()
	if err := h.ClubRepo.Update(club); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, club)
}

package handler

import (
	"encoding/json"
	"gofootball/app/model"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAllPlayers get all players
func GetAllPlayers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	player := []model.Player{}
	db.Find(&player)
	respondJSON(w, http.StatusOK, player)
}

// CreatePlayer insert new player
func CreatePlayer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	player := model.Player{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&player); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&player).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, player)
}

// GetPlayer by player name
func GetPlayer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	player := getPlayerOr404(db, name, w, r)
	if player == nil {
		return
	}
	respondJSON(w, http.StatusOK, player)
}

// UpdatePlayer by player name
func UpdatePlayer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	player := getPlayerOr404(db, name, w, r)
	if player == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&player); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&player).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, player)
}

// DeletePlayer by player name
func DeletePlayer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	player := getPlayerOr404(db, name, w, r)
	if player == nil {
		return
	}
	if err := db.Delete(&player).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// DisablePlayer by player name
func DisablePlayer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	player := getPlayerOr404(db, name, w, r)
	if player == nil {
		return
	}
	player.Disable()
	if err := db.Save(&player).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, player)
}

// EnablePlayer by player name
func EnablePlayer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	player := getClubOr404(db, name, w, r)
	if player == nil {
		return
	}
	player.Enable()
	if err := db.Save(&player).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, player)
}

// getPlayerOr404 gets a player instance if exists, or respond the 404 error otherwise
func getPlayerOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.Player {
	player := model.Player{}
	if err := db.First(&player, model.Player{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &player
}

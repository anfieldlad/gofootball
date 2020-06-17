package handler

import (
	"encoding/json"
	"gofootball/app/model"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAllClubs get all clubs
func GetAllClubs(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	club := []model.Club{}
	db.Find(&club)
	respondJSON(w, http.StatusOK, club)
}

// CreateClub insert new club
func CreateClub(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	club := model.Club{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&club); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&club).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, club)
}

// GetClub by club name
func GetClub(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	club := getClubOr404(db, name, w, r)
	if club == nil {
		return
	}
	respondJSON(w, http.StatusOK, club)
}

// UpdateClub by club name
func UpdateClub(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	club := getClubOr404(db, name, w, r)
	if club == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&club); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&club).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, club)
}

// DeleteClub by club name
func DeleteClub(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	club := getClubOr404(db, name, w, r)
	if club == nil {
		return
	}
	if err := db.Delete(&club).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// DisableClub by club name
func DisableClub(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	club := getClubOr404(db, name, w, r)
	if club == nil {
		return
	}
	club.Disable()
	if err := db.Save(&club).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, club)
}

// EnableClub by club name
func EnableClub(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	club := getClubOr404(db, name, w, r)
	if club == nil {
		return
	}
	club.Enable()
	if err := db.Save(&club).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, club)
}

// getClubOr404 gets a club instance if exists, or respond the 404 error otherwise
func getClubOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.Club {
	club := model.Club{}
	if err := db.First(&club, model.Club{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &club
}

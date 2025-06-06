package handler

import "gofootball/internal/repository"

// Handler groups all http handlers.
type Handler struct {
	ClubRepo   repository.ClubRepository
	PlayerRepo repository.PlayerRepository
}

// New creates a new Handler with given repositories.
func New(clubRepo repository.ClubRepository, playerRepo repository.PlayerRepository) *Handler {
	return &Handler{ClubRepo: clubRepo, PlayerRepo: playerRepo}
}

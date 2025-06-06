package server

import (
	"fmt"
	"log"
	"net/http"

	"gofootball/internal/config"
	"gofootball/internal/handler"
	"gofootball/internal/model"
	"gofootball/internal/repository"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Server has router and db instances
type Server struct {
	Router  *mux.Router
	DB      *gorm.DB
	Handler *handler.Handler
}

// Initialize with predefined configuration
func (s *Server) Initialize(cfg *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Server,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.Charset)

	db, err := gorm.Open(cfg.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	s.DB = model.DBMigrate(db)
	//model.Seed(db)

	clubRepo := &repository.GormClubRepository{DB: s.DB}
	playerRepo := &repository.GormPlayerRepository{DB: s.DB}
	s.Handler = handler.New(clubRepo, playerRepo)

	s.Router = mux.NewRouter()
	s.setRouters()
}

// Set all required routers
func (s *Server) setRouters() {
	// Routing for handling the projects
	s.Get("/clubs", s.Handler.GetAllClubs)
	s.Post("/club", s.Handler.CreateClub)
	s.Get("/club/{name}", s.Handler.GetClub)
	s.Put("/club/{name}", s.Handler.UpdateClub)
	s.Delete("/club/{name}", s.Handler.DeleteClub)
	s.Put("/club/{name}/disable", s.Handler.DisableClub)
	s.Put("/club/{name}/enable", s.Handler.EnableClub)

	s.Get("/players", s.Handler.GetAllPlayers)
	s.Post("/player", s.Handler.CreatePlayer)
	s.Get("/player/{name}", s.Handler.GetPlayer)
	s.Put("/player/{name}", s.Handler.UpdatePlayer)
	s.Delete("/player/{name}", s.Handler.DeletePlayer)
	s.Put("/player/{name}/disable", s.Handler.DisablePlayer)
	s.Put("/player/{name}/enable", s.Handler.EnablePlayer)
}

// Get : Wrap the router for GET method
func (s *Server) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("GET")
}

// Post : Wrap the router for POST method
func (s *Server) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("POST")
}

// Put : Wrap the router for PUT method
func (s *Server) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete : Wrap the router for DELETE method
func (s *Server) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app
func (s *Server) Run(host string) {
	log.Println("Starting the server at http://127.0.0.1" + host + "/")
	log.Println("Quit the server with CONTROL-C.")
	log.Fatal(http.ListenAndServe(host, s.Router))
}

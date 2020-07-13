package app

import (
	"fmt"
	"log"
	"net/http"

	"gofootball/app/handler"
	"gofootball/app/model"
	"gofootball/config"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Server,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	//model.Seed(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/clubs", a.GetAllClubs)
	a.Post("/club", a.CreateClub)
	a.Get("/club/{name}", a.GetClub)
	a.Put("/club/{name}", a.UpdateClub)
	a.Delete("/club/{name}", a.DeleteClub)
	a.Put("/club/{name}/disable", a.DisableClub)
	a.Put("/club/{name}/enable", a.EnableClub)

	a.Get("/players", a.GetAllPlayers)
	a.Post("/player", a.CreatePlayer)
	a.Get("/player/{name}", a.GetPlayer)
	a.Put("/player/{name}", a.UpdatePlayer)
	a.Delete("/player/{name}", a.DeletePlayer)
	a.Put("/player/{name}/disable", a.DisablePlayer)
	a.Put("/player/{name}/enable", a.EnablePlayer)
}

// Get : Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post : Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put : Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete : Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// GetAllClubs : Handlers
func (a *App) GetAllClubs(w http.ResponseWriter, r *http.Request) {
	handler.GetAllClubs(a.DB, w, r)
}

// CreateClub : Handlers
func (a *App) CreateClub(w http.ResponseWriter, r *http.Request) {
	handler.CreateClub(a.DB, w, r)
}

// GetClub : Handlers
func (a *App) GetClub(w http.ResponseWriter, r *http.Request) {
	handler.GetClub(a.DB, w, r)
}

// UpdateClub : Handlers
func (a *App) UpdateClub(w http.ResponseWriter, r *http.Request) {
	handler.UpdateClub(a.DB, w, r)
}

// DeleteClub : Handlers
func (a *App) DeleteClub(w http.ResponseWriter, r *http.Request) {
	handler.DeleteClub(a.DB, w, r)
}

// DisableClub : Handlers
func (a *App) DisableClub(w http.ResponseWriter, r *http.Request) {
	handler.DisableClub(a.DB, w, r)
}

// EnableClub : Handlers
func (a *App) EnableClub(w http.ResponseWriter, r *http.Request) {
	handler.EnableClub(a.DB, w, r)
}

// GetAllPlayers : Handlers
func (a *App) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllPlayers(a.DB, w, r)
}

// CreatePlayer : Handlers
func (a *App) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	handler.CreatePlayer(a.DB, w, r)
}

// GetPlayer : Handlers
func (a *App) GetPlayer(w http.ResponseWriter, r *http.Request) {
	handler.GetPlayer(a.DB, w, r)
}

// UpdatePlayer : Handlers
func (a *App) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	handler.UpdatePlayer(a.DB, w, r)
}

// DeletePlayer : Handlers
func (a *App) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	handler.DeletePlayer(a.DB, w, r)
}

// DisablePlayer : Handlers
func (a *App) DisablePlayer(w http.ResponseWriter, r *http.Request) {
	handler.DisablePlayer(a.DB, w, r)
}

// EnablePlayer : Handlers
func (a *App) EnablePlayer(w http.ResponseWriter, r *http.Request) {
	handler.EnablePlayer(a.DB, w, r)
}

// Run the app
func (a *App) Run(host string) {
	log.Println("Starting the server at http://127.0.0.1" + host + "/")
	log.Println("Quit the server with CONTROL-C.")
	log.Fatal(http.ListenAndServe(host, a.Router))
}

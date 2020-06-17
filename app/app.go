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
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/clubs", a.GetAllClubs)
	a.Post("/clubs", a.CreateClub)
	a.Get("/clubs/{name}", a.GetClub)
	a.Put("/clubs/{name}", a.UpdateClub)
	a.Delete("/clubs/{name}", a.DeleteClub)
	a.Put("/clubs/{name}/disable", a.DisableClub)
	a.Put("/clubs/{name}/enable", a.EnableClub)
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

// Run the app
func (a *App) Run(host string) {
	log.Println("Starting the server at http://127.0.0.1" + host + "/")
	log.Println("Quit the server with CONTROL-C.")
	log.Fatal(http.ListenAndServe(host, a.Router))
}

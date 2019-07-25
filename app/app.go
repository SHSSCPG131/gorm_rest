package app
import (
	"fmt"
	"ghorm_candidate_backened _code/app/handler"
	"ghorm_candidate_backened _code/config"
	"ghorm_candidate_backened _code/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)
// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}
// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		main.Username,
		main.Password,
		main.Name,
		main.Charset)

	db, err := gorm.Open(main.Dialect, dbURI)
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
	a.Get("/getdetail", a.GetAllDetails)
	a.Post("/createuser", a.CreateUser)
	a.Get("/getuserbyemail/{title}", a.GetUserByEmail)
	a.Put("/updatedetail/{title}", a.UpdateDetail)
	a.Delete("/deletedetail/{title}", a.DeleteDetail)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage users Data
func (a *App) GetAllDetails(w http.ResponseWriter, r *http.Request) {
	main.GetAllDetails(a.DB, w, r)
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	main.CreateUser(a.DB, w, r)
}

func (a *App) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	main.GetDetail(a.DB, w, r)
}

func (a *App) UpdateDetail(w http.ResponseWriter, r *http.Request) {
	main.UpdateDetail(a.DB, w, r)
}

func (a *App) DeleteDetail(w http.ResponseWriter, r *http.Request) {
	main.Deletedetail(a.DB, w, r)
}
// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

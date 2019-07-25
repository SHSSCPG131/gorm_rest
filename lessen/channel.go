package lessen

import (
"encoding/json"
	"log"
"net/http"
"os"
"github.com/gorilla/mux"
"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
"github.com/rs/cors"
_ "github.com/jinzhu/gorm/dialects/postgres"
)

type new_details struct {
	gorm.Model
	Id			int
	Name           string
	Source         string
	Phone_number   string
	Experience      string
	Ctc            string
	Ectc           string
	Np             string
	Status         string
	Interview_date string
	Email          string       //required
	Applied_for    string  //required`
}
var db *gorm.DB
var err error
func main() {
	router := mux.NewRouter()
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=dvdrental password=12345")
	defer db.Close()
if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&new_details{})
	router.HandleFunc("/getdetail", GetResources).Methods("GET")
	router.HandleFunc("/getdetail/{email}", GetResource).Methods("GET")
	router.HandleFunc("/createdetail", CreateResource).Methods("POST")
	router.HandleFunc("/deletedetail/{id}", DeleteResource).Methods("DELETE")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}

func GetResources(w http.ResponseWriter, r *http.Request) {
	var resources []new_details
	db.Find(&resources)
	json.NewEncoder(w).Encode(&resources)
}

func GetResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var resource new_details
	db.First(&resource, params["email"])
	json.NewEncoder(w).Encode(&resource)
}

func CreateResource(w http.ResponseWriter, r *http.Request) {
	var resource new_details
	json.NewDecoder(r.Body).Decode(&resource)
	db.Create(&resource)
	json.NewEncoder(w).Encode(&resource)
}
func DeleteResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var resource new_details
	db.First(&resource, params["id"])
	db.Delete(&resource)
	var resources []new_details
	db.Find(&resources)
	json.NewEncoder(w).Encode(&resources)
}
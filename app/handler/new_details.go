package main
import (
	"encoding/json"
	"ghorm_candidate_backened _code/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)
//new_details[]>>database,,new_detail>>>query
func GetAllDetails(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	new_detail_show := []model.new_details{}
	db.Find(&new_detail_show)
	respondJSON(w, http.StatusOK, new_detail_show)
}
func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	new_detail_user := model.new_details{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&new_detail_user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := db.Save(&new_detail_user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, new_detail_user)
}
func GetDetail(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	new_detail := getnew_detailOr404(db, email, w, r)
	if new_detail == nil {
		return
	}
	respondJSON(w, http.StatusOK, new_detail)
}
func UpdateDetail(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	new_detail := getnew_detailOr404(db, name, w, r)
	if new_detail == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&new_detail); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&new_detail).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, new_detail)
}
func Deletedetail(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	new_detail := getnew_detailOr404(db, name, w, r)
	if new_detail == nil {
		return
	}
	if err := db.Delete(&new_detail).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}
// getnew_detailOr404 gets a detail instance if exists, or respond the 404 error otherwise
func getnew_detailOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.new_details {
	new_detail := model.new_details{}
	if err := db.First(&new_detail, model.new_details{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &new_detail
}


package model
import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
//used database dvdrental and table name is new_details
type new_details struct {
	gorm.Model
	Id			int `gorm:"unique" json:"id"`
	Name           string `json:"name"`
	Source         string `json:"source"`
	Phone_number   string `json:"phone_number"`
	Experience      string `json:"experience"`
	Ctc            string `json:"ctc"`
	Ectc           string `json:"ectc"`
	Np             string `json:"np"`
	Status         bool `json:"status"`
	Interview_date string `json:"interview_date"`
	Email          string `json:"email"`       //required
	Applied_for    string `json:"applied_for"` //required`

}
// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&new_details{})
	return db
}
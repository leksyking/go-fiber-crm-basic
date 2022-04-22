package lead

import "github.com/jinzhu/gorm"

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLead() {

}

func GetLeads() {

}

func NewLead() {

}

func DeleteLead() {

}

package models

type Subject struct {
	Subject_Id          uint   `gorm:"primaryKey;autoIncrement" json:"subject_id"`
	Subject_Description string `gorm:"not null" json:"subject_description"`
	Is_active           bool   `gorm:"default:true" json:"is_active"`
	Module_id           uint   `json:"module_id"`
}

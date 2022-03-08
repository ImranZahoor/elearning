package models

type Module struct {
	Module_Id          uint      `gorm:"primaryKey;autoIncrement" json:"module_id"`
	Module_Description string    `gorm:"not null" json:"module_description"`
	Is_active          bool      `gorm:"default:true" json:"is_active"`
	Subjects           []Subject `gorm:"foreignKey:Module_id" json:"subjects"`
}

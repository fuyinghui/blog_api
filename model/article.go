package model

type Article struct {
	ID        uint       `gorm:"primary_key"`
	Title     string     `gorm:"primary_key;type:varchar(255);not null;"`
	Content   string     `gorm:"type:text;not null"`
	Status    int        `gorm:"type:int(1);default:0"`
	CreatedAt TimeNormal `gorm:"column:created_at;default:null"`
	UpdatedAt TimeNormal `gorm:"column:updated_at;default:null"`
}

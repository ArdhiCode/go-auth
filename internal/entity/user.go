package entity

import "github.com/google/uuid"

type User struct {
	Timestamp
	Id       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name     string    `gorm:"type:varchar(255);not null" json:"name"`
	Email    string    `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password string    `gorm:"type:text;not null" json:"password"`
	Role     string    `gorm:"type:varchar(50);not null" json:"role"`
}

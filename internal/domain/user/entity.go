package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	ADMIN  UserRole = "ADMIN"
	DRIVER UserRole = "DRIVER"
)

type User struct {
	Id        uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null"`
	Email     string         `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password  string         `json:"-" gorm:"type:text;not null"`
	Role      UserRole       `json:"role" gorm:"type:user_role;default:'DRIVER';not null"`
	Phone     string         `json:"phone" gorm:"type:varchar(20)"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"type:timestamp;index"`
}

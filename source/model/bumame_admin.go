package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	// Id field
	Id uint64 `json:"id" gorm:"autoIncrement"`

	// Main field
	AdminUUID   uuid.UUID `json:"admin_uuid"`
	Name        string    `json:"name"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Role        string    `json:"role"`                       // tenaga-kesehatan, dokter, pengawas, admin, b2b-sales, b2b-ops
	Position    string    `json:"position"`                   // manager, head, staff
	Source      string    `json:"source" gorm:"default:null"` // internal, external
	IsConfirmed int       `json:"is_confirmed"`
	Status      int       `json:"status"`
	PhotoFile   string    `json:"photo_file" gorm:"default:null"`
	Longitude   float64   `json:"longitude" gorm:"default:null"`
	Latitude    float64   `json:"latitude" gorm:"default:null"`
	SkillIds    string    `json:"skill_ids" gorm:"default:null"`
	AdminType   string    `json:"admin_type" gorm:"default:null"` // perawat, analis

	CurrentLocationJson string `json:"current_location_json" gorm:"default:null"` // longitude, latitude, address

	IsDeleted uint   `json:"is_deleted" gorm:"default:0"`
	Password  string `json:"-"`

	// Encrypted field
	EncryptedPhone []byte `json:"encrypted_phone" gorm:"default:null"`

	// Timestamp
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (model *Admin) BeforeCreate(tx *gorm.DB) error {
	model.AdminUUID = uuid.New()
	return nil
}

func (Admin) TableName() string {
	return "bumame_admin"
}

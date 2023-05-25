package model

import "github.com/google/uuid"

type User struct {
	BaseModel
	ID       uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Wallet   []Wallet  `json:"wallet"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

package model

import "github.com/google/uuid"

type Wallet struct {
	BaseModel
	Address string    `json:"address" gorm:"primaryKey"`
	Name    string    `json:"name"`
	UserID  uuid.UUID `json:"user_id"`
	User    User      `json:"user" gorm:"foreignKey:user_id;references:id"`
}

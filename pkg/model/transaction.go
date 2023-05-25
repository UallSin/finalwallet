package model

import "github.com/google/uuid"

type Transaction struct {
	BaseModel
	ID uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`

	FromAddress string `json:"from_address"`
	WalletFrom  Wallet `json:"wallet_from" gorm:"foreignKey:from_address;references:address"`

	ToAddress string `json:"to_address"`
	WalletTo  Wallet `json:"wallet_to" gorm:"foreignKey:to_address;references:address"`

	TokenAddress string  `json:"token_address"`
	Token        Token   `json:"token" gorm:"foreignKey:token_address;references:address"`
	Amount       float64 `json:"amount"`
}

// TODO: transaction
// tạo transaction: A -> 100 ETH -> B
// B -> 10 ETH ->

package models

import "time"

type User struct {
	ID        uint      `json:"-" gorm:"primaryKey;autoincrement"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Address   string    `json:"address"`
	Balance   float64   `json:"balance" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

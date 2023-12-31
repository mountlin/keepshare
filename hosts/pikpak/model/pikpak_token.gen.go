// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameToken = "pikpak_token"

// Token mapped from table <pikpak_token>
type Token struct {
	UserID       string    `gorm:"column:user_id;primaryKey" json:"user_id"`
	AccessToken  string    `gorm:"column:access_token;not null" json:"access_token"`
	RefreshToken string    `gorm:"column:refresh_token;not null" json:"refresh_token"`
	Expiration   time.Time `gorm:"column:expiration;not null" json:"expiration"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName Token's table name
func (*Token) TableName() string {
	return TableNameToken
}

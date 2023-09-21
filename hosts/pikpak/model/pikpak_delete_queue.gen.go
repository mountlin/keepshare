// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDeleteQueue = "pikpak_delete_queue"

// DeleteQueue mapped from table <pikpak_delete_queue>
type DeleteQueue struct {
	WorkerUserID     string    `gorm:"column:worker_user_id;primaryKey" json:"worker_user_id"`
	OriginalLinkHash string    `gorm:"column:original_link_hash;primaryKey" json:"original_link_hash"`
	Status           string    `gorm:"column:status;not null" json:"status"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	NextTrigger      time.Time `gorm:"column:next_trigger;not null;default:CURRENT_TIMESTAMP" json:"next_trigger"`
	Ext              string    `gorm:"column:ext" json:"ext"`
}

// TableName DeleteQueue's table name
func (*DeleteQueue) TableName() string {
	return TableNameDeleteQueue
}

package models

import "time"

type Absensi struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	AbsensiType string    `gorm:"type:enum('masuk', 'keluar');not null" json:"absensi_type"`
	CreatedAt   time.Time `json:"created_at"`
	User        User
}

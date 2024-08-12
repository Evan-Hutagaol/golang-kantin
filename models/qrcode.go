package models

import "time"

type QRCode struct {
    ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
    Code        string    `gorm:"type:varchar(255);unique;not null" json:"code"`
    AbsensiType string    `gorm:"type:enum('masuk', 'keluar');not null" json:"absensi_type"`
    ValidFrom   time.Time `json:"valid_from"`
    ValidTo     time.Time `json:"valid_to"`
    GeneratedAt time.Time `json:"generated_at"`
}
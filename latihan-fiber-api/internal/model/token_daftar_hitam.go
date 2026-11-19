package model

import "time"

type TokenDaftarHitam struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Token       string    `gorm:"size:512;uniqueIndex" json:"-"`
	Kedaluwarsa time.Time `json:"kedaluwarsa"`
	CreatedAt   time.Time `json:"dibuat_pada"`
}

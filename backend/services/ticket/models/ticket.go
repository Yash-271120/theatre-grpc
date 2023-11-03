package models

import (
	"time"
)

type Ticket struct {
	Id        uint32    `json:"id" gorm:"primary_key"`
	HallId    uint32    `json:"hall_id"`
	Seats     uint32    `json:"seats"`
	EntryCode string    `json:"entry_code"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}

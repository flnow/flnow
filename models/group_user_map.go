package models

import "time"

type GroupUserMap struct{
	ID    int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Group int `gorm:"Column:group" json:"group"`
	User int `gorm:"Column:user" json:"user"`
	CreatedAt time.Time `gorm:"Column:createdAt" json:"createdAt"`
}
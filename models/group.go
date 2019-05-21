package models

import "time"

type Group struct {
	ID        int       `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"Column:name;NOT NULL" json:"name"`
	Avatar    string    `gorm:"Column:avatar" json:"avatar"`
	Owner     int       `gorm:"Column:owner" json:"owner"`
	CreatedAt time.Time `gorm:"Column:createdAt" json:"createdAt"`
}

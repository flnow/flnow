package models

import "time"

// GroupUserMap is a mapping of group and user
type GroupUserMap struct {
	ID        int       `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Group     int       `gorm:"Column:group" json:"group"`
	User      int       `gorm:"Column:user" json:"user"`
	CreatedAt time.Time `gorm:"Column:createdAt" json:"createdAt"`
}

// TableName of GroupUserMap entity for ORM
func (GroupUserMap) TableName() string {
	return "groupUserMap"
}

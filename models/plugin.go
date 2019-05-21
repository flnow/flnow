package models

import "time"

type Plugin struct {
	ID             int       `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Name           string    `gorm:"Column:name;NOT NULL" json:"name"`
	Desc           string    `gorm:"Column:desc;Size:32768" json:"desc"`
	ContainerImage string    `gorm:"Column:containerImage;NOT NULL" json:"containerImage"`
	Owner          int       `gorm:"Column:owner;NOT NULL" json:"owner"`
	CreatedAt      time.Time `gorm:"Column:createdAt" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"Column:updatedAt" json:"updatedAt"`
}

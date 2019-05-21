package models

import "time"

type User struct {
	ID          int       `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	DisplayName string    `gorm:"Column:displayName;NOT NULL" json:"displayName"`
	UserName    string    `gorm:"Column:userName;NOT NULL" json:"userName"`
	Password    string    `gorm:"Column:password;NOT NULL;Size:32768" json:"-"`
	Groups      string    `gorm:"Column:groups;Size:32768" json:"groups"`
	Email       string    `gorm:"Column:email" json:"email"`
	Phone       string    `gorm:"Column:phone" json:"phone"`
	Avatar      string    `gorm:"Column:avatar" json:"avatar"`
	CreatedAt   time.Time `gorm:"Column:createdAt" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"Column:updatedAt" json:"updatedAt"`
}

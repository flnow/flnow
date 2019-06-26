package auth

import "time"

// User entity
type User struct {
	ID          int       `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
	DisplayName string    `gorm:"Column:displayName;NOT NULL" json:"displayName" form:"displayName"`
	UserName    string    `gorm:"Column:userName;NOT NULL" json:"username" form:"username"`
	Password    string    `gorm:"Column:password;NOT NULL;type:text" json:"-" form:"password"`
	Groups      string    `gorm:"Column:groups;type:text" json:"groups" form:"groups"`
	Email       string    `gorm:"Column:email" json:"email" form:"email"`
	Phone       string    `gorm:"Column:phone" json:"phone" form:"phone"`
	Avatar      string    `gorm:"Column:avatar" json:"avatar" form:"avatar"`
	CreatedAt   time.Time `gorm:"Column:createdAt" json:"createdAt" form:"createdAt"`
	UpdatedAt   time.Time `gorm:"Column:updatedAt" json:"updatedAt" form:"updatedAt"`
}

// TableName of User entity for ORM
func (User) TableName() string {
	return "user"
}

func (u *User) initialize() {
	initTime := time.Now()
	u.CreatedAt = initTime
	u.UpdatedAt = initTime
}

package models

// GlobalConfiguration entity
// To store global configurations with different types
type GlobalConfiguration struct {
	ID    int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Owner int    `gorm:"Column:owner;NOT NULL" json:"owner"`
	Type  string `gorm:"Column:type;NOT NULL" json:"type"`
	Key   string `gorm:"Column:key;NOT NULL" json:"key"`
	Value string `gorm:"Column:value;NOT NULL" json:"value"`
}

// TableName of GlobalConfiguration entity for ORM
func (GlobalConfiguration) TableName() string {
	return "globalConfiguration"
}

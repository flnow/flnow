package models

import "time"

// Node entity
type Node struct {
	ID           int       `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	FlowID       int       `gorm:"Column:flowId" json:"flowId"`
	ParentID     int       `gorm:"Column:parentId" json:"parentId"`
	PluginID     int       `gorm:"Column:pluginId" json:"pluginId"`
	Sequence     int       `gorm:"Column:sequence" json:"sequence"`
	RunCondition string    `gorm:"Column:runCondition" json:"runCondition"`
	CreatedAt    time.Time `gorm:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"updatedAt" json:"updatedAt"`
}

// TableName of Node entity for ORM
func (Node) TableName() string {
	return "node"
}

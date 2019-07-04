package node

import (
	"time"
)

// Node entity
type Node struct {
	ID               int       `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	FlowID           int       `gorm:"Column:flowId" json:"flowId"`
	ParentID         int       `gorm:"Column:parentId" json:"parentId"`
	PluginID         int       `gorm:"Column:pluginId" json:"pluginId"`
	Sequence         int       `gorm:"Column:sequence" json:"sequence"`
	PreviewCondition string    `gorm:"Column:previewCondition" json:"previewCondition"` // preview node run condition
	CreatedAt        time.Time `gorm:"createdAt" json:"createdAt"`
	UpdatedAt        time.Time `gorm:"updatedAt" json:"updatedAt"`
}

// TableName of Node entity for ORM
func (Node) TableName() string {
	return "node"
}

func (n *Node) initialize() {
	initTime := time.Now()
	n.CreatedAt = initTime
	n.UpdatedAt = initTime
	if n.ParentID < 0 {
		// Root node
		n.PreviewCondition = "ANY"
		n.Sequence = 1
		n.ParentID = -1
	}
}

package models

type NodeConfiguration struct {
	ID     int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	NodeID int    `gorm:"Column:nodeId;NOT NULL" json:"nodeId"`
	Key    string `gorm:"Column:key;NOT NULL" json:"key"`
	Value  string `gorm:"Column:value;NOT NULL" json:"value"`
}

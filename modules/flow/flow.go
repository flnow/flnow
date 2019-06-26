package flow

import (
	"time"
)

// Flow entity
type Flow struct {
	ID                  int       `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
	Name                string    `gorm:"Column:name;NOT NULL" json:"name" form:"name"`
	TriggerType         string    `gorm:"Column:triggerType;NOT NULL" json:"triggerType" form:"triggerType"`
	CronExpression      string    `gorm:"Column:cronExpression;type:varchar(50)" json:"cron" form:"cron"`
	CronTimeZone        string    `gorm:"Column:cronTimeZone;type:varchar(50)" json:"cronTimeZone" form:"cronTimeZone"`
	State               string    `gorm:"Column:state;size:50" json:"state" form:"state"`
	Token               string    `gorm:"Column:token;size:255" json:"-" form:"token"`
	NodeCount           int       `gorm:"Column:nodeCount" json:"nodeCount" form:"nodeCount"`
	LastExecutedAt      time.Time `gorm:"Column:lastExecutedAt" json:"lastExecutedAt" form:"lastExecutedAt"`
	LastExecutedSummary string    `gorm:"Column:lastExecutedSummary;type:text" json:"lastExecutedSummary" form:"lastExecutedSummary"`
	RunAt               string    `gorm:"Column:runAt;type:text" json:"runAt" form:"runAt"` // 在哪个worker上被执行
	HostedOn            string    `gorm:"Column:hostedOn" json:"hostedOn" form:"hostedOn"`  // 在哪里触发，slave
	Pointer             string    `gorm:"Column:pointer" json:"-"`                          // 创建后的cronJob在内存里的ID

	Owner int    `gorm:"Column:owner" json:"owner"`
	Tags  string `gorm:"Column:tags;type:text" json:"tags"`

	Comment   string    `gorm:"Column:comment;type:text" json:"comment"`
	CreatedAt time.Time `gorm:"Column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"Column:updatedAt" json:"updatedAt"`
}

// TableName of Flow entity for ORM
func (Flow) TableName() string {
	return "flow"
}

// initialize the flow fields which need be initialized
func (f *Flow) initialize() {
	initTime := time.Now()
	f.Owner = 1
	f.State = "CREATED"
	f.NodeCount = 0
	f.RunAt = "ANY"
	f.Pointer = "-1"
	f.LastExecutedAt = initTime
	f.CreatedAt = initTime
	f.UpdatedAt = initTime
}

package flow

import (
	"time"

	"github.com/google/uuid"
)

// Flow entity
type Flow struct {
	ID                  string    `gorm:"Column:id;PRIMARY_KEY" json:"id" form:"id"`
	Name                string    `gorm:"Column:name;NOT NULL" json:"name" form:"name"`
	TriggerType         string    `gorm:"Column:triggerType;NOT NULL" json:"triggerType" form:"triggerType"`
	CronExpression      string    `gorm:"Column:cronExpression;type:varchar(50)" json:"cron" form:"cron"`
	CronTimeZone        string    `gorm:"Column:cronTimeZone;type:varchar(50)" json:"cronTimeZone" form:"cronTimeZone"`
	State               string    `gorm:"Column:state;size:50" json:"state" form:"state"`
	Token               string    `gorm:"Column:token;size:255" json:"-" form:"token"`
	NodeCount           int       `gorm:"Column:nodeCount" json:"nodeCount" form:"nodeCount"`
	LastExecutedAt      time.Time `gorm:"Column:lastExecutedAt" json:"lastExecutedAt" form:"lastExecutedAt"`
	LastExecutedSummary string    `gorm:"Column:lastExecutedSummary;type:text" json:"lastExecutedSummary" form:"lastExecutedSummary"`
	RunAt               string    `gorm:"Column:runAt;type:text" json:"runAt" form:"runAt"`          // 在哪个worker上被执行
	HostedOn            string    `gorm:"Column:hostedOn" json:"hostedOn" form:"hostedOn"`           // 在哪里触发，slave
	Pointer             string    `gorm:"Column:pointer" json:"-"`                                   // 创建后的cronJob在内存里的ID
	Pipeline            string    `gorm:"Column:pipeline;type:text" json:"pipeline" form:"pipeline"` // 具体配置项

	Owner int    `gorm:"Column:owner" json:"owner"`
	Tags  string `gorm:"Column:tags;type:text" json:"tags" form:"tags"`

	Comment   string    `gorm:"Column:comment;type:text" json:"comment"`
	CreatedAt time.Time `gorm:"Column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"Column:updatedAt" json:"updatedAt"`
}

// TableName of Flow entity for ORM
func (Flow) TableName() string {
	return "flow"
}

// Node entity
type Node struct {
	ID           string `json:"id" gorm:"column:id;PRIMARY_KEY"`
	FlowID       string `json:"flowID" gorm:"column:flowID"`
	ParentID     string `json:"parentID" gorm:"column:parentID"`
	Plugin       string `json:"plugin" gorm:"column:plugin"`
	Sequence     int    `json:"sequence" gorm:"column:sequence"`
	RunCondition string `json:"runCondition" gorm:"column:runCondition"` // run condition based parent's executed output.
}

// TableName of Node entity for ORM
func (Node) TableName() string {
	return "node"
}

// NodeConfiguration is the configs of one Node
type NodeConfiguration struct {
	ID     string `json:"id" gorm:"column:id;PRIMARY_KEY"`
	NodeID string `json:"nodeID" gorm:"column:nodeID;PRIMARY_KEY"`
	Key    string `json:"k" gorm:"column:k"`
	Value  string `json:"v" gorm:"column:v"`
}

// TableName of NodeConfiguration entity for ORM
func (NodeConfiguration) TableName() string {
	return "nodeConfiguration"
}

// Pipeline struct of Node+Configuration
type Pipeline struct {
	Plugin        string            `json:"plugin"`
	RunCondition  string            `json:"runCondition"`
	Configuration map[string]string `json:"configuration"`
	Success       *Pipeline         `json:"success"` // Another pipeline node
	Failure       *Pipeline         `json:"failure"` // Another pipeline node
	Any           *Pipeline         `json:"any"`     // Another pipeline node
}

// IsZero method to check the pipeline instance zero value or not
func (p *Pipeline) IsZero() bool {
	if p != nil && p.Plugin == "" {
		return true
	}
	return false
}

// ToRelational method to convert Pipeline instance to Node and NodeConfiguration
func (p *Pipeline) ToRelational(flow, parent, condition string, sequence int) (nodes []Node, configs []NodeConfiguration) {
	if len(parent) == 0 {
		// root/entry node
		condition = "ANY"
	}
	currNode := Node{
		ID:           uuid.New().String(),
		FlowID:       flow,
		ParentID:     parent,
		Plugin:       p.Plugin,
		Sequence:     sequence,
		RunCondition: condition,
	}

	currConfigs := []NodeConfiguration{}

	for confK, confV := range p.Configuration {
		currConfigs = append(currConfigs, NodeConfiguration{
			ID:     uuid.New().String(),
			NodeID: currNode.ID,
			Key:    confK,
			Value:  confV,
		})
	}

	nodes = append(nodes, currNode)
	configs = append(configs, currConfigs...)

	if p.Success != nil && !p.Success.IsZero() {
		successNodes, successConfigs := p.Success.ToRelational(flow, currNode.ID, "SUCCESS", currNode.Sequence+1)
		nodes = append(nodes, successNodes...)
		configs = append(configs, successConfigs...)
	}
	if p.Failure != nil && !p.Failure.IsZero() {
		failureNodes, failureConfigs := p.Failure.ToRelational(flow, currNode.ID, "FAILURE", currNode.Sequence+1)
		nodes = append(nodes, failureNodes...)
		configs = append(configs, failureConfigs...)
	}
	if p.Any != nil && !p.Any.IsZero() {
		anyNodes, anyConfigs := p.Any.ToRelational(flow, currNode.ID, "ANY", currNode.Sequence+1)
		nodes = append(nodes, anyNodes...)
		configs = append(configs, anyConfigs...)
	}

	return
}

// initialize the flow fields which need be initialized
func (f *Flow) initialize() {
	initTime := time.Now()
	if len(f.ID) == 0 {
		f.ID = uuid.New().String()
	}
	f.Owner = 1
	f.State = "CREATED"
	f.NodeCount = 0
	f.RunAt = "ANY"
	f.Pointer = "-1"
	f.LastExecutedAt = initTime
	f.CreatedAt = initTime
	f.UpdatedAt = initTime
}

package defination

// TaskChannel 任务触发通道
type TaskChannel struct {
	Name    string
	Channel chan ServiceCommand
}

// ServiceCommand 各个服务之间的命令传递
type ServiceCommand struct {
	ID      string
	Cmd     string
	Args    string
	Channel chan bool
}

// TaskContentClone 克隆任务
type TaskContentClone struct {
	Address     string `json:"address"`
	Destination string `json:"destination"`
}

// WebServiceReturn WebServiceReturn
type WebServiceReturn struct {
	Code  int
	Error string
}

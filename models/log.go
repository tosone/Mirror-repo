package models

import (
	"github.com/jinzhu/gorm"
)

// Log ..
type Log struct {
	gorm.Model
	RepoID uint   // 外键
	Cmd    string // 任务名称
	Status string // 是否成功
	Msg    string // 消息
}

// Create ..
func (log *Log) Create() error {
	return engine.Create(log).Error
}

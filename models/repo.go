package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tosone/mirrorepo/common/defination"
)

// Repo ..
type Repo struct {
	gorm.Model
	Address         string                // 仓库地址
	Name            string                // 本地存储所用的名字
	AliasName       string                // 真正存储的地方
	Travel          int                   // 仓库两次更新之间的时间间隔
	LastTraveled    time.Time             // 仓库上次被更新的时间
	Status          defination.RepoStatus // 仓库状态
	CommitCount     uint64                // commit 数量
	LastCommitCount uint64                // 之前 commit 数量
	HistoryInfoID   string                // 外键 ID
}

// Create ..
func (repo *Repo) Create() error {
	return engine.Create(repo).Error
}

// Delete ..
func (repo *Repo) Delete() error {
	return engine.Delete(&repo).Error
}

// Find ..
func (repo *Repo) Find() (r *Repo, err error) {
	err = engine.Where(repo.ID).Find(r).Error
	return
}

// UpdateByID ..
func (repo *Repo) UpdateByID() error {
	return engine.Model(new(Repo)).Where(repo.ID).Updates(repo).Error
}

// GetAll ..
func (repo *Repo) GetAll() (repos *([]Repo), err error) {
	repos = new([]Repo)
	err = engine.Find(repos).Error
	return
}

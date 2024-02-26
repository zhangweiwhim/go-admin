package models

import (
	"go-admin/common/models"
)

type TbTerm struct {
	models.Model

	Name      string `json:"name" gorm:"type:varchar(128);comment:名称"`
	StartTime string `json:"startTime" gorm:"type:varchar(128);comment:开始时间"`
	EndTime   string `json:"endTime" gorm:"type:varchar(128);comment:结束时间"`
	models.ModelTime
	models.ControlBy
}

func (TbTerm) TableName() string {
	return "tb_term"
}

func (e *TbTerm) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbTerm) GetId() interface{} {
	return e.Id
}

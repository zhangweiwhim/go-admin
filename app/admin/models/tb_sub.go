package models

import (
	"go-admin/common/models"
)

type TbSub struct {
	models.Model

	Name string `json:"name" gorm:"type:varchar(128);comment:名称"`
	models.ModelTime
	models.ControlBy
}

func (TbSub) TableName() string {
	return "tb_sub"
}

func (e *TbSub) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbSub) GetId() interface{} {
	return e.Id
}

package models

import (
	"go-admin/common/models"
)

type TbTeachAll struct {
	models.Model

	Grade    string `json:"grade" gorm:"type:varchar(100);comment:任教年级"`
	Class    string `json:"class" gorm:"type:varchar(100);comment:任教班级"`
	Subject  string `json:"subject" gorm:"type:varchar(100);comment:任教学科"`
	TechName string `json:"techName" gorm:"type:varchar(100);comment:教师名称"`
	TechNo   string `json:"techNo" gorm:"type:varchar(100);comment:教师工号"`
	TechTerm string `json:"techTerm" gorm:"type:varchar(100);comment:任教学期"`
	models.ModelTime
	models.ControlBy
}

func (TbTeachAll) TableName() string {
	return "tb_teach_all"
}

func (e *TbTeachAll) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbTeachAll) GetId() interface{} {
	return e.Id
}

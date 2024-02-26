package models

import (
	"time"

	"go-admin/common/models"
)

type TbStudent struct {
	models.Model

	Name        string    `json:"name" gorm:"type:varchar(128);comment:名称"`
	Grade       string    `json:"grade" gorm:"type:varchar(100);comment:年级"`
	Class       string    `json:"class" gorm:"type:varchar(100);comment:班级"`
	HeadTeacher string    `json:"headTeacher" gorm:"type:varchar(100);comment:班主任"`
	StuNo       string    `json:"stuNo" gorm:"type:varchar(100);comment:学号"`
	Sex         string    `json:"sex" gorm:"type:int;comment:性别"`
	Age         string    `json:"age" gorm:"type:int;comment:年龄"`
	IdNo        string    `json:"idNo" gorm:"type:varchar(100);comment:身份证"`
	Address     string    `json:"address" gorm:"type:varchar(100);comment:家庭地址"`
	Tel         string    `json:"tel" gorm:"type:varchar(100);comment:电话"`
	ParentsName string    `json:"parentsName" gorm:"type:varchar(100);comment:家长姓名"`
	ParentsTel  string    `json:"parentsTel" gorm:"type:varchar(100);comment:家长电话"`
	InTime      time.Time `json:"inTime" gorm:"type:timestamp;comment:入学时间"`
	models.ModelTime
	models.ControlBy
}

func (TbStudent) TableName() string {
	return "tb_student"
}

func (e *TbStudent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbStudent) GetId() interface{} {
	return e.Id
}

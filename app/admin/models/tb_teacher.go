package models

import (
    "time"

	"go-admin/common/models"

)

type TbTeacher struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(100);comment:姓名"` 
    TecNo string `json:"tecNo" gorm:"type:varchar(100);comment:工号"` 
    TecNowSub string `json:"tecNowSub" gorm:"type:varchar(100);comment:任职学科"` 
    TecNowGrade string `json:"tecNowGrade" gorm:"type:varchar(100);comment:任职年级"` 
    TecNowClass string `json:"tecNowClass" gorm:"type:varchar(100);comment:任职班级"` 
    Sex string `json:"sex" gorm:"type:int;comment:性别"` 
    InTime time.Time `json:"inTime" gorm:"type:timestamp;comment:入值时间"` 
    Age string `json:"age" gorm:"type:int;comment:年级"` 
    IdNo string `json:"idNo" gorm:"type:varchar(100);comment:身份证号"` 
    Address string `json:"address" gorm:"type:varchar(100);comment:地址"` 
    Tel string `json:"tel" gorm:"type:varchar(100);comment:电话"` 
    Honorary string `json:"honorary" gorm:"type:varchar(100);comment:荣誉"` 
    Resume string `json:"resume" gorm:"type:varchar(100);comment:简介"` 
    OtherTel string `json:"otherTel" gorm:"type:varchar(100);comment:其他联系方式"` 
    models.ModelTime
    models.ControlBy
}

func (TbTeacher) TableName() string {
    return "tb_teacher"
}

func (e *TbTeacher) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbTeacher) GetId() interface{} {
	return e.Id
}
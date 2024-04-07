package models

import (

	"go-admin/common/models"

)

type TbChengji struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(128);comment:名称"` 
    Grade string `json:"grade" gorm:"type:varchar(100);comment:年级"` 
    Class string `json:"class" gorm:"type:varchar(100);comment:班级"` 
    Xueqi string `json:"xueqi" gorm:"type:varchar(100);comment:学期"` 
    KaoNo string `json:"kaoNo" gorm:"type:varchar(100);comment:考号"` 
    YuwenScore string `json:"yuwenScore" gorm:"type:varchar(100);comment:语文"` 
    ShuxueScore string `json:"shuxueScore" gorm:"type:varchar(100);comment:数学"` 
    YingyuScore string `json:"yingyuScore" gorm:"type:varchar(100);comment:英语"` 
    ZhengzhiScore string `json:"zhengzhiScore" gorm:"type:varchar(100);comment:政治"` 
    LishiScore string `json:"lishiScore" gorm:"type:varchar(100);comment:历史"` 
    DiliScore string `json:"diliScore" gorm:"type:varchar(100);comment:地理"` 
    WuliScore string `json:"wuliScore" gorm:"type:varchar(100);comment:物理"` 
    HuaxueScore string `json:"huaxueScore" gorm:"type:varchar(100);comment:化学"` 
    ShengwuScore string `json:"shengwuScore" gorm:"type:varchar(100);comment:生物"` 
    KaoshiTime string `json:"kaoshiTime" gorm:"type:varchar(100);comment:考试时间"` 
    models.ModelTime
    models.ControlBy
}

func (TbChengji) TableName() string {
    return "tb_chengji"
}

func (e *TbChengji) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbChengji) GetId() interface{} {
	return e.Id
}
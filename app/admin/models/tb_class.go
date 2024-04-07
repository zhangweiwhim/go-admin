package models

import (

	"go-admin/common/models"

)

type TbClass struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(128);comment:名称"` 
    Xueqi string `json:"xueqi" gorm:"type:varchar(100);comment:学期"` 
    Grade string `json:"grade" gorm:"type:varchar(100);comment:年级"` 
    Class string `json:"class" gorm:"type:varchar(100);comment:班级"` 
    HeadTeacher string `json:"headTeacher" gorm:"type:varchar(100);comment:班主任"` 
    YuwenTeacher string `json:"yuwenTeacher" gorm:"type:varchar(100);comment:语文老师"` 
    ShuxueTeacher string `json:"shuxueTeacher" gorm:"type:varchar(100);comment:数学老师"` 
    YingyuTeacher string `json:"yingyuTeacher" gorm:"type:varchar(100);comment:英语老师"` 
    ZhengzhiTeacher string `json:"zhengzhiTeacher" gorm:"type:varchar(100);comment:政治老师"` 
    LishiTeacher string `json:"lishiTeacher" gorm:"type:varchar(100);comment:历史老师"` 
    DiliTeacher string `json:"diliTeacher" gorm:"type:varchar(100);comment:地理老师"` 
    WuliTeacher string `json:"wuliTeacher" gorm:"type:varchar(100);comment:物理老师"` 
    HuaxueTeacher string `json:"huaxueTeacher" gorm:"type:varchar(100);comment:化学老师"` 
    ShengwuTeacher string `json:"shengwuTeacher" gorm:"type:varchar(100);comment:生物老师"` 
    Nums string `json:"nums" gorm:"type:int;comment:人数"` 
    models.ModelTime
    models.ControlBy
}

func (TbClass) TableName() string {
    return "tb_class"
}

func (e *TbClass) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbClass) GetId() interface{} {
	return e.Id
}
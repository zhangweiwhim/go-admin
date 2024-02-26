package models

import (
	"go-admin/common/models"
)

type TbExam struct {
	models.Model

	ExamGrade string `json:"examGrade" gorm:"type:varchar(100);comment:年级"`
	ExamClass string `json:"examClass" gorm:"type:varchar(100);comment:班级"`
	ExamStart string `json:"examStart" gorm:"type:varchar(100);comment:开始时间"`
	ExamEnd   string `json:"examEnd" gorm:"type:varchar(100);comment:结束时间"`
	ExamType  string `json:"examType" gorm:"type:int;comment:考试批次"`
	ExamSub   string `json:"examSub" gorm:"type:int;comment:考试科目"`
	TermName  string `json:"termName" gorm:"type:varchar(100);comment:学期名称"`
	models.ModelTime
	models.ControlBy
}

func (TbExam) TableName() string {
	return "tb_exam"
}

func (e *TbExam) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TbExam) GetId() interface{} {
	return e.Id
}

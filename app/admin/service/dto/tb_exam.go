package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TbExamGetPageReq struct {
	dto.Pagination `search:"-"`
	ExamStart      string `form:"examStart"  search:"type:contains;column:exam_start;table:tb_exam" comment:"开始时间"`
	ExamType       string `form:"examType"  search:"type:contains;column:exam_type;table:tb_exam" comment:"考试批次"`
	TbExamOrder
}

type TbExamOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:tb_exam"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_exam"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_exam"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_exam"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_exam"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_exam"`
	ExamGrade string `form:"examGradeOrder"  search:"type:order;column:exam_grade;table:tb_exam"`
	ExamClass string `form:"examClassOrder"  search:"type:order;column:exam_class;table:tb_exam"`
	ExamStart string `form:"examStartOrder"  search:"type:order;column:exam_start;table:tb_exam"`
	ExamEnd   string `form:"examEndOrder"  search:"type:order;column:exam_end;table:tb_exam"`
	ExamType  string `form:"examTypeOrder"  search:"type:order;column:exam_type;table:tb_exam"`
	ExamSub   string `form:"examSubOrder"  search:"type:order;column:exam_sub;table:tb_exam"`
	TermName  string `form:"termNameOrder"  search:"type:order;column:term_name;table:tb_exam"`
}

func (m *TbExamGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbExamInsertReq struct {
	Id        int    `json:"-" comment:"主键编码"` // 主键编码
	ExamGrade string `json:"examGrade" comment:"年级"`
	ExamClass string `json:"examClass" comment:"班级"`
	ExamStart string `json:"examStart" comment:"开始时间"`
	ExamEnd   string `json:"examEnd" comment:"结束时间"`
	ExamType  string `json:"examType" comment:"考试批次"`
	ExamSub   string `json:"examSub" comment:"考试科目"`
	TermName  string `json:"termName" comment:"学期名称"`
	common.ControlBy
}

func (s *TbExamInsertReq) Generate(model *models.TbExam) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.ExamGrade = s.ExamGrade
	model.ExamClass = s.ExamClass
	model.ExamStart = s.ExamStart
	model.ExamEnd = s.ExamEnd
	model.ExamType = s.ExamType
	model.ExamSub = s.ExamSub
	model.TermName = s.TermName
}

func (s *TbExamInsertReq) GetId() interface{} {
	return s.Id
}

type TbExamUpdateReq struct {
	Id        int    `uri:"id" comment:"主键编码"` // 主键编码
	ExamGrade string `json:"examGrade" comment:"年级"`
	ExamClass string `json:"examClass" comment:"班级"`
	ExamStart string `json:"examStart" comment:"开始时间"`
	ExamEnd   string `json:"examEnd" comment:"结束时间"`
	ExamType  string `json:"examType" comment:"考试批次"`
	ExamSub   string `json:"examSub" comment:"考试科目"`
	TermName  string `json:"termName" comment:"学期名称"`
	common.ControlBy
}

func (s *TbExamUpdateReq) Generate(model *models.TbExam) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.ExamGrade = s.ExamGrade
	model.ExamClass = s.ExamClass
	model.ExamStart = s.ExamStart
	model.ExamEnd = s.ExamEnd
	model.ExamType = s.ExamType
	model.ExamSub = s.ExamSub
	model.TermName = s.TermName
}

func (s *TbExamUpdateReq) GetId() interface{} {
	return s.Id
}

// TbExamGetReq 功能获取请求参数
type TbExamGetReq struct {
	Id int `uri:"id"`
}

func (s *TbExamGetReq) GetId() interface{} {
	return s.Id
}

// TbExamDeleteReq 功能删除请求参数
type TbExamDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbExamDeleteReq) GetId() interface{} {
	return s.Ids
}

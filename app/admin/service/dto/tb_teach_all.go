package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TbTeachAllGetPageReq struct {
	dto.Pagination `search:"-"`
	Grade          string `form:"grade"  search:"type:contains;column:grade;table:tb_teach_all" comment:"任教年级"`
	Class          string `form:"class"  search:"type:contains;column:class;table:tb_teach_all" comment:"任教班级"`
	Subject        string `form:"subject"  search:"type:contains;column:subject;table:tb_teach_all" comment:"任教学科"`
	TechName       string `form:"techName"  search:"type:contains;column:tech_name;table:tb_teach_all" comment:"教师名称"`
	TechNo         string `form:"techNo"  search:"type:contains;column:tech_no;table:tb_teach_all" comment:"教师工号"`
	TechTerm       string `form:"techTerm"  search:"type:contains;column:tech_term;table:tb_teach_all" comment:"任教学期"`
	TbTeachAllOrder
}

type TbTeachAllOrder struct {
	Grade    string `form:"gradeOrder"  search:"type:order;column:grade;table:tb_teach_all"`
	Class    string `form:"classOrder"  search:"type:order;column:class;table:tb_teach_all"`
	Subject  string `form:"subjectOrder"  search:"type:order;column:subject;table:tb_teach_all"`
	TechName string `form:"techNameOrder"  search:"type:order;column:tech_name;table:tb_teach_all"`
	TechNo   string `form:"techNoOrder"  search:"type:order;column:tech_no;table:tb_teach_all"`
	TechTerm string `form:"techTermOrder"  search:"type:order;column:tech_term;table:tb_teach_all"`
	Id       string `form:"idOrder"  search:"type:order;column:id;table:tb_teach_all"`
}

func (m *TbTeachAllGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbTeachAllInsertReq struct {
	Grade    string `json:"grade" comment:"任教年级"`
	Class    string `json:"class" comment:"任教班级"`
	Subject  string `json:"subject" comment:"任教学科"`
	TechName string `json:"techName" comment:"教师名称"`
	TechNo   string `json:"techNo" comment:"教师工号"`
	TechTerm string `json:"techTerm" comment:"任教学期"`
	Id       int    `json:"-" comment:""` //
	common.ControlBy
}

func (s *TbTeachAllInsertReq) Generate(model *models.TbTeachAll) {
	model.Grade = s.Grade
	model.Class = s.Class
	model.Subject = s.Subject
	model.TechName = s.TechName
	model.TechNo = s.TechNo
	model.TechTerm = s.TechTerm
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
}

func (s *TbTeachAllInsertReq) GetId() interface{} {
	return s.Id
}

type TbTeachAllUpdateReq struct {
	Grade    string `json:"grade" comment:"任教年级"`
	Class    string `json:"class" comment:"任教班级"`
	Subject  string `json:"subject" comment:"任教学科"`
	TechName string `json:"techName" comment:"教师名称"`
	TechNo   string `json:"techNo" comment:"教师工号"`
	TechTerm string `json:"techTerm" comment:"任教学期"`
	Id       int    `uri:"id" comment:""` //
	common.ControlBy
}

func (s *TbTeachAllUpdateReq) Generate(model *models.TbTeachAll) {
	model.Grade = s.Grade
	model.Class = s.Class
	model.Subject = s.Subject
	model.TechName = s.TechName
	model.TechNo = s.TechNo
	model.TechTerm = s.TechTerm
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
}

func (s *TbTeachAllUpdateReq) GetId() interface{} {
	return s.Id
}

// TbTeachAllGetReq 功能获取请求参数
type TbTeachAllGetReq struct {
	Id int `uri:"id"`
}

func (s *TbTeachAllGetReq) GetId() interface{} {
	return s.Id
}

// TbTeachAllDeleteReq 功能删除请求参数
type TbTeachAllDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbTeachAllDeleteReq) GetId() interface{} {
	return s.Ids
}

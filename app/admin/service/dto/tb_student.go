package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TbStudentGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:tb_student" comment:"名称"`
	Grade          string `form:"grade"  search:"type:contains;column:grade;table:tb_student" comment:"年级"`
	Class          string `form:"class"  search:"type:contains;column:class;table:tb_student" comment:"班级"`
	StuNo          string `form:"stuNo"  search:"type:contains;column:stu_no;table:tb_student" comment:"学号"`
	TbStudentOrder
}

type TbStudentOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:tb_student"`
	Name        string `form:"nameOrder"  search:"type:order;column:name;table:tb_student"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_student"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_student"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_student"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_student"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_student"`
	Grade       string `form:"gradeOrder"  search:"type:order;column:grade;table:tb_student"`
	Class       string `form:"classOrder"  search:"type:order;column:class;table:tb_student"`
	HeadTeacher string `form:"headTeacherOrder"  search:"type:order;column:head_teacher;table:tb_student"`
	StuNo       string `form:"stuNoOrder"  search:"type:order;column:stu_no;table:tb_student"`
	Sex         string `form:"sexOrder"  search:"type:order;column:sex;table:tb_student"`
	Age         string `form:"ageOrder"  search:"type:order;column:age;table:tb_student"`
	IdNo        string `form:"idNoOrder"  search:"type:order;column:id_no;table:tb_student"`
	Address     string `form:"addressOrder"  search:"type:order;column:address;table:tb_student"`
	Tel         string `form:"telOrder"  search:"type:order;column:tel;table:tb_student"`
	ParentsName string `form:"parentsNameOrder"  search:"type:order;column:parents_name;table:tb_student"`
	ParentsTel  string `form:"parentsTelOrder"  search:"type:order;column:parents_tel;table:tb_student"`
	InTime      string `form:"inTimeOrder"  search:"type:order;column:in_time;table:tb_student"`
}

func (m *TbStudentGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbStudentInsertReq struct {
	Id          int       `json:"-" comment:"主键编码"` // 主键编码
	Name        string    `json:"name" comment:"名称"`
	Grade       string    `json:"grade" comment:"年级"`
	Class       string    `json:"class" comment:"班级"`
	HeadTeacher string    `json:"headTeacher" comment:"班主任"`
	StuNo       string    `json:"stuNo" comment:"学号"`
	Sex         string    `json:"sex" comment:"性别"`
	Age         string    `json:"age" comment:"年龄"`
	IdNo        string    `json:"idNo" comment:"身份证"`
	Address     string    `json:"address" comment:"家庭地址"`
	Tel         string    `json:"tel" comment:"电话"`
	ParentsName string    `json:"parentsName" comment:"家长姓名"`
	ParentsTel  string    `json:"parentsTel" comment:"家长电话"`
	InTime      time.Time `json:"inTime" comment:"入学时间"`
	common.ControlBy
}

func (s *TbStudentInsertReq) Generate(model *models.TbStudent) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Grade = s.Grade
	model.Class = s.Class
	model.HeadTeacher = s.HeadTeacher
	model.StuNo = s.StuNo
	model.Sex = s.Sex
	model.Age = s.Age
	model.IdNo = s.IdNo
	model.Address = s.Address
	model.Tel = s.Tel
	model.ParentsName = s.ParentsName
	model.ParentsTel = s.ParentsTel
	model.InTime = s.InTime
}

func (s *TbStudentInsertReq) GetId() interface{} {
	return s.Id
}

type TbStudentUpdateReq struct {
	Id          int       `uri:"id" comment:"主键编码"` // 主键编码
	Name        string    `json:"name" comment:"名称"`
	Grade       string    `json:"grade" comment:"年级"`
	Class       string    `json:"class" comment:"班级"`
	HeadTeacher string    `json:"headTeacher" comment:"班主任"`
	StuNo       string    `json:"stuNo" comment:"学号"`
	Sex         string    `json:"sex" comment:"性别"`
	Age         string    `json:"age" comment:"年龄"`
	IdNo        string    `json:"idNo" comment:"身份证"`
	Address     string    `json:"address" comment:"家庭地址"`
	Tel         string    `json:"tel" comment:"电话"`
	ParentsName string    `json:"parentsName" comment:"家长姓名"`
	ParentsTel  string    `json:"parentsTel" comment:"家长电话"`
	InTime      time.Time `json:"inTime" comment:"入学时间"`
	common.ControlBy
}

func (s *TbStudentUpdateReq) Generate(model *models.TbStudent) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Grade = s.Grade
	model.Class = s.Class
	model.HeadTeacher = s.HeadTeacher
	model.StuNo = s.StuNo
	model.Sex = s.Sex
	model.Age = s.Age
	model.IdNo = s.IdNo
	model.Address = s.Address
	model.Tel = s.Tel
	model.ParentsName = s.ParentsName
	model.ParentsTel = s.ParentsTel
	model.InTime = s.InTime
}

func (s *TbStudentUpdateReq) GetId() interface{} {
	return s.Id
}

// TbStudentGetReq 功能获取请求参数
type TbStudentGetReq struct {
	Id int `uri:"id"`
}

func (s *TbStudentGetReq) GetId() interface{} {
	return s.Id
}

// TbStudentDeleteReq 功能删除请求参数
type TbStudentDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbStudentDeleteReq) GetId() interface{} {
	return s.Ids
}

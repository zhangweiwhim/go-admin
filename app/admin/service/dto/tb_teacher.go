package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TbTeacherGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:contains;column:name;table:tb_teacher" comment:"姓名"`
    TecNo string `form:"tecNo"  search:"type:exact;column:tec_no;table:tb_teacher" comment:"工号"`
    TecNowSub string `form:"tecNowSub"  search:"type:exact;column:tec_now_sub;table:tb_teacher" comment:"任职学科"`
    TecNowGrade string `form:"tecNowGrade"  search:"type:exact;column:tec_now_grade;table:tb_teacher" comment:"任职年级"`
    TecNowClass string `form:"tecNowClass"  search:"type:exact;column:tec_now_class;table:tb_teacher" comment:"任职班级"`
    TbTeacherOrder
}

type TbTeacherOrder struct {
    Name string `form:"nameOrder"  search:"type:order;column:name;table:tb_teacher"`
    TecNo string `form:"tecNoOrder"  search:"type:order;column:tec_no;table:tb_teacher"`
    TecNowSub string `form:"tecNowSubOrder"  search:"type:order;column:tec_now_sub;table:tb_teacher"`
    TecNowGrade string `form:"tecNowGradeOrder"  search:"type:order;column:tec_now_grade;table:tb_teacher"`
    TecNowClass string `form:"tecNowClassOrder"  search:"type:order;column:tec_now_class;table:tb_teacher"`
    Sex string `form:"sexOrder"  search:"type:order;column:sex;table:tb_teacher"`
    InTime string `form:"inTimeOrder"  search:"type:order;column:in_time;table:tb_teacher"`
    Age string `form:"ageOrder"  search:"type:order;column:age;table:tb_teacher"`
    IdNo string `form:"idNoOrder"  search:"type:order;column:id_no;table:tb_teacher"`
    Address string `form:"addressOrder"  search:"type:order;column:address;table:tb_teacher"`
    Tel string `form:"telOrder"  search:"type:order;column:tel;table:tb_teacher"`
    Honorary string `form:"honoraryOrder"  search:"type:order;column:honorary;table:tb_teacher"`
    Resume string `form:"resumeOrder"  search:"type:order;column:resume;table:tb_teacher"`
    OtherTel string `form:"otherTelOrder"  search:"type:order;column:other_tel;table:tb_teacher"`
    Id string `form:"idOrder"  search:"type:order;column:id;table:tb_teacher"`
    
}

func (m *TbTeacherGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbTeacherInsertReq struct {
    Name string `json:"name" comment:"姓名"`
    TecNo string `json:"tecNo" comment:"工号"`
    TecNowSub string `json:"tecNowSub" comment:"任职学科"`
    TecNowGrade string `json:"tecNowGrade" comment:"任职年级"`
    TecNowClass string `json:"tecNowClass" comment:"任职班级"`
    Sex string `json:"sex" comment:"性别"`
    InTime time.Time `json:"inTime" comment:"入值时间"`
    Age string `json:"age" comment:"年级"`
    IdNo string `json:"idNo" comment:"身份证号"`
    Address string `json:"address" comment:"地址"`
    Tel string `json:"tel" comment:"电话"`
    Honorary string `json:"honorary" comment:"荣誉"`
    Resume string `json:"resume" comment:"简介"`
    OtherTel string `json:"otherTel" comment:"其他联系方式"`
    Id int `json:"-" comment:""` // 
    common.ControlBy
}

func (s *TbTeacherInsertReq) Generate(model *models.TbTeacher)  {
    model.Name = s.Name
    model.TecNo = s.TecNo
    model.TecNowSub = s.TecNowSub
    model.TecNowGrade = s.TecNowGrade
    model.TecNowClass = s.TecNowClass
    model.Sex = s.Sex
    model.InTime = s.InTime
    model.Age = s.Age
    model.IdNo = s.IdNo
    model.Address = s.Address
    model.Tel = s.Tel
    model.Honorary = s.Honorary
    model.Resume = s.Resume
    model.OtherTel = s.OtherTel
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
}

func (s *TbTeacherInsertReq) GetId() interface{} {
	return s.Id
}

type TbTeacherUpdateReq struct {
    Name string `json:"name" comment:"姓名"`
    TecNo string `json:"tecNo" comment:"工号"`
    TecNowSub string `json:"tecNowSub" comment:"任职学科"`
    TecNowGrade string `json:"tecNowGrade" comment:"任职年级"`
    TecNowClass string `json:"tecNowClass" comment:"任职班级"`
    Sex string `json:"sex" comment:"性别"`
    InTime time.Time `json:"inTime" comment:"入值时间"`
    Age string `json:"age" comment:"年级"`
    IdNo string `json:"idNo" comment:"身份证号"`
    Address string `json:"address" comment:"地址"`
    Tel string `json:"tel" comment:"电话"`
    Honorary string `json:"honorary" comment:"荣誉"`
    Resume string `json:"resume" comment:"简介"`
    OtherTel string `json:"otherTel" comment:"其他联系方式"`
    Id int `uri:"id" comment:""` // 
    common.ControlBy
}

func (s *TbTeacherUpdateReq) Generate(model *models.TbTeacher)  {
    model.Name = s.Name
    model.TecNo = s.TecNo
    model.TecNowSub = s.TecNowSub
    model.TecNowGrade = s.TecNowGrade
    model.TecNowClass = s.TecNowClass
    model.Sex = s.Sex
    model.InTime = s.InTime
    model.Age = s.Age
    model.IdNo = s.IdNo
    model.Address = s.Address
    model.Tel = s.Tel
    model.Honorary = s.Honorary
    model.Resume = s.Resume
    model.OtherTel = s.OtherTel
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
}

func (s *TbTeacherUpdateReq) GetId() interface{} {
	return s.Id
}

// TbTeacherGetReq 功能获取请求参数
type TbTeacherGetReq struct {
     Id int `uri:"id"`
}
func (s *TbTeacherGetReq) GetId() interface{} {
	return s.Id
}

// TbTeacherDeleteReq 功能删除请求参数
type TbTeacherDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbTeacherDeleteReq) GetId() interface{} {
	return s.Ids
}

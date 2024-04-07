package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TbClassGetPageReq struct {
	dto.Pagination     `search:"-"`
    Xueqi string `form:"xueqi"  search:"type:exact;column:xueqi;table:tb_class" comment:"学期"`
    Grade string `form:"grade"  search:"type:contains;column:grade;table:tb_class" comment:"年级"`
    Class string `form:"class"  search:"type:contains;column:class;table:tb_class" comment:"班级"`
    HeadTeacher string `form:"headTeacher"  search:"type:contains;column:head_teacher;table:tb_class" comment:"班主任"`
    TbClassOrder
}

type TbClassOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:tb_class"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:tb_class"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_class"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_class"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_class"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_class"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_class"`
    Xueqi string `form:"xueqiOrder"  search:"type:order;column:xueqi;table:tb_class"`
    Grade string `form:"gradeOrder"  search:"type:order;column:grade;table:tb_class"`
    Class string `form:"classOrder"  search:"type:order;column:class;table:tb_class"`
    HeadTeacher string `form:"headTeacherOrder"  search:"type:order;column:head_teacher;table:tb_class"`
    YuwenTeacher string `form:"yuwenTeacherOrder"  search:"type:order;column:yuwen_teacher;table:tb_class"`
    ShuxueTeacher string `form:"shuxueTeacherOrder"  search:"type:order;column:shuxue_teacher;table:tb_class"`
    YingyuTeacher string `form:"yingyuTeacherOrder"  search:"type:order;column:yingyu_teacher;table:tb_class"`
    ZhengzhiTeacher string `form:"zhengzhiTeacherOrder"  search:"type:order;column:zhengzhi_teacher;table:tb_class"`
    LishiTeacher string `form:"lishiTeacherOrder"  search:"type:order;column:lishi_teacher;table:tb_class"`
    DiliTeacher string `form:"diliTeacherOrder"  search:"type:order;column:dili_teacher;table:tb_class"`
    WuliTeacher string `form:"wuliTeacherOrder"  search:"type:order;column:wuli_teacher;table:tb_class"`
    HuaxueTeacher string `form:"huaxueTeacherOrder"  search:"type:order;column:huaxue_teacher;table:tb_class"`
    ShengwuTeacher string `form:"shengwuTeacherOrder"  search:"type:order;column:shengwu_teacher;table:tb_class"`
    Nums string `form:"numsOrder"  search:"type:order;column:nums;table:tb_class"`
    
}

func (m *TbClassGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbClassInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    Name string `json:"name" comment:"名称"`
    Xueqi string `json:"xueqi" comment:"学期"`
    Grade string `json:"grade" comment:"年级"`
    Class string `json:"class" comment:"班级"`
    HeadTeacher string `json:"headTeacher" comment:"班主任"`
    YuwenTeacher string `json:"yuwenTeacher" comment:"语文老师"`
    ShuxueTeacher string `json:"shuxueTeacher" comment:"数学老师"`
    YingyuTeacher string `json:"yingyuTeacher" comment:"英语老师"`
    ZhengzhiTeacher string `json:"zhengzhiTeacher" comment:"政治老师"`
    LishiTeacher string `json:"lishiTeacher" comment:"历史老师"`
    DiliTeacher string `json:"diliTeacher" comment:"地理老师"`
    WuliTeacher string `json:"wuliTeacher" comment:"物理老师"`
    HuaxueTeacher string `json:"huaxueTeacher" comment:"化学老师"`
    ShengwuTeacher string `json:"shengwuTeacher" comment:"生物老师"`
    Nums string `json:"nums" comment:"人数"`
    common.ControlBy
}

func (s *TbClassInsertReq) Generate(model *models.TbClass)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
    model.Xueqi = s.Xueqi
    model.Grade = s.Grade
    model.Class = s.Class
    model.HeadTeacher = s.HeadTeacher
    model.YuwenTeacher = s.YuwenTeacher
    model.ShuxueTeacher = s.ShuxueTeacher
    model.YingyuTeacher = s.YingyuTeacher
    model.ZhengzhiTeacher = s.ZhengzhiTeacher
    model.LishiTeacher = s.LishiTeacher
    model.DiliTeacher = s.DiliTeacher
    model.WuliTeacher = s.WuliTeacher
    model.HuaxueTeacher = s.HuaxueTeacher
    model.ShengwuTeacher = s.ShengwuTeacher
    model.Nums = s.Nums
}

func (s *TbClassInsertReq) GetId() interface{} {
	return s.Id
}

type TbClassUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    Name string `json:"name" comment:"名称"`
    Xueqi string `json:"xueqi" comment:"学期"`
    Grade string `json:"grade" comment:"年级"`
    Class string `json:"class" comment:"班级"`
    HeadTeacher string `json:"headTeacher" comment:"班主任"`
    YuwenTeacher string `json:"yuwenTeacher" comment:"语文老师"`
    ShuxueTeacher string `json:"shuxueTeacher" comment:"数学老师"`
    YingyuTeacher string `json:"yingyuTeacher" comment:"英语老师"`
    ZhengzhiTeacher string `json:"zhengzhiTeacher" comment:"政治老师"`
    LishiTeacher string `json:"lishiTeacher" comment:"历史老师"`
    DiliTeacher string `json:"diliTeacher" comment:"地理老师"`
    WuliTeacher string `json:"wuliTeacher" comment:"物理老师"`
    HuaxueTeacher string `json:"huaxueTeacher" comment:"化学老师"`
    ShengwuTeacher string `json:"shengwuTeacher" comment:"生物老师"`
    Nums string `json:"nums" comment:"人数"`
    common.ControlBy
}

func (s *TbClassUpdateReq) Generate(model *models.TbClass)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
    model.Xueqi = s.Xueqi
    model.Grade = s.Grade
    model.Class = s.Class
    model.HeadTeacher = s.HeadTeacher
    model.YuwenTeacher = s.YuwenTeacher
    model.ShuxueTeacher = s.ShuxueTeacher
    model.YingyuTeacher = s.YingyuTeacher
    model.ZhengzhiTeacher = s.ZhengzhiTeacher
    model.LishiTeacher = s.LishiTeacher
    model.DiliTeacher = s.DiliTeacher
    model.WuliTeacher = s.WuliTeacher
    model.HuaxueTeacher = s.HuaxueTeacher
    model.ShengwuTeacher = s.ShengwuTeacher
    model.Nums = s.Nums
}

func (s *TbClassUpdateReq) GetId() interface{} {
	return s.Id
}

// TbClassGetReq 功能获取请求参数
type TbClassGetReq struct {
     Id int `uri:"id"`
}
func (s *TbClassGetReq) GetId() interface{} {
	return s.Id
}

// TbClassDeleteReq 功能删除请求参数
type TbClassDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbClassDeleteReq) GetId() interface{} {
	return s.Ids
}

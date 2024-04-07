package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TbChengjiGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:exact;column:name;table:tb_chengji" comment:"名称"`
    Grade string `form:"grade"  search:"type:exact;column:grade;table:tb_chengji" comment:"年级"`
    KaoNo string `form:"kaoNo"  search:"type:exact;column:kao_no;table:tb_chengji" comment:"考号"`
    KaoshiTime string `form:"kaoshiTime"  search:"type:exact;column:kaoshi_time;table:tb_chengji" comment:"考试时间"`
    TbChengjiOrder
}

type TbChengjiOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:tb_chengji"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:tb_chengji"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_chengji"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_chengji"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_chengji"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_chengji"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_chengji"`
    Grade string `form:"gradeOrder"  search:"type:order;column:grade;table:tb_chengji"`
    Class string `form:"classOrder"  search:"type:order;column:class;table:tb_chengji"`
    Xueqi string `form:"xueqiOrder"  search:"type:order;column:xueqi;table:tb_chengji"`
    KaoNo string `form:"kaoNoOrder"  search:"type:order;column:kao_no;table:tb_chengji"`
    YuwenScore string `form:"yuwenScoreOrder"  search:"type:order;column:yuwen_score;table:tb_chengji"`
    ShuxueScore string `form:"shuxueScoreOrder"  search:"type:order;column:shuxue_score;table:tb_chengji"`
    YingyuScore string `form:"yingyuScoreOrder"  search:"type:order;column:yingyu_score;table:tb_chengji"`
    ZhengzhiScore string `form:"zhengzhiScoreOrder"  search:"type:order;column:zhengzhi_score;table:tb_chengji"`
    LishiScore string `form:"lishiScoreOrder"  search:"type:order;column:lishi_score;table:tb_chengji"`
    DiliScore string `form:"diliScoreOrder"  search:"type:order;column:dili_score;table:tb_chengji"`
    WuliScore string `form:"wuliScoreOrder"  search:"type:order;column:wuli_score;table:tb_chengji"`
    HuaxueScore string `form:"huaxueScoreOrder"  search:"type:order;column:huaxue_score;table:tb_chengji"`
    ShengwuScore string `form:"shengwuScoreOrder"  search:"type:order;column:shengwu_score;table:tb_chengji"`
    KaoshiTime string `form:"kaoshiTimeOrder"  search:"type:order;column:kaoshi_time;table:tb_chengji"`
    
}

func (m *TbChengjiGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbChengjiInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    Name string `json:"name" comment:"名称"`
    Grade string `json:"grade" comment:"年级"`
    Class string `json:"class" comment:"班级"`
    Xueqi string `json:"xueqi" comment:"学期"`
    KaoNo string `json:"kaoNo" comment:"考号"`
    YuwenScore string `json:"yuwenScore" comment:"语文"`
    ShuxueScore string `json:"shuxueScore" comment:"数学"`
    YingyuScore string `json:"yingyuScore" comment:"英语"`
    ZhengzhiScore string `json:"zhengzhiScore" comment:"政治"`
    LishiScore string `json:"lishiScore" comment:"历史"`
    DiliScore string `json:"diliScore" comment:"地理"`
    WuliScore string `json:"wuliScore" comment:"物理"`
    HuaxueScore string `json:"huaxueScore" comment:"化学"`
    ShengwuScore string `json:"shengwuScore" comment:"生物"`
    KaoshiTime string `json:"kaoshiTime" comment:"考试时间"`
    common.ControlBy
}

func (s *TbChengjiInsertReq) Generate(model *models.TbChengji)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
    model.Grade = s.Grade
    model.Class = s.Class
    model.Xueqi = s.Xueqi
    model.KaoNo = s.KaoNo
    model.YuwenScore = s.YuwenScore
    model.ShuxueScore = s.ShuxueScore
    model.YingyuScore = s.YingyuScore
    model.ZhengzhiScore = s.ZhengzhiScore
    model.LishiScore = s.LishiScore
    model.DiliScore = s.DiliScore
    model.WuliScore = s.WuliScore
    model.HuaxueScore = s.HuaxueScore
    model.ShengwuScore = s.ShengwuScore
    model.KaoshiTime = s.KaoshiTime
}

func (s *TbChengjiInsertReq) GetId() interface{} {
	return s.Id
}

type TbChengjiUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    Name string `json:"name" comment:"名称"`
    Grade string `json:"grade" comment:"年级"`
    Class string `json:"class" comment:"班级"`
    Xueqi string `json:"xueqi" comment:"学期"`
    KaoNo string `json:"kaoNo" comment:"考号"`
    YuwenScore string `json:"yuwenScore" comment:"语文"`
    ShuxueScore string `json:"shuxueScore" comment:"数学"`
    YingyuScore string `json:"yingyuScore" comment:"英语"`
    ZhengzhiScore string `json:"zhengzhiScore" comment:"政治"`
    LishiScore string `json:"lishiScore" comment:"历史"`
    DiliScore string `json:"diliScore" comment:"地理"`
    WuliScore string `json:"wuliScore" comment:"物理"`
    HuaxueScore string `json:"huaxueScore" comment:"化学"`
    ShengwuScore string `json:"shengwuScore" comment:"生物"`
    KaoshiTime string `json:"kaoshiTime" comment:"考试时间"`
    common.ControlBy
}

func (s *TbChengjiUpdateReq) Generate(model *models.TbChengji)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
    model.Grade = s.Grade
    model.Class = s.Class
    model.Xueqi = s.Xueqi
    model.KaoNo = s.KaoNo
    model.YuwenScore = s.YuwenScore
    model.ShuxueScore = s.ShuxueScore
    model.YingyuScore = s.YingyuScore
    model.ZhengzhiScore = s.ZhengzhiScore
    model.LishiScore = s.LishiScore
    model.DiliScore = s.DiliScore
    model.WuliScore = s.WuliScore
    model.HuaxueScore = s.HuaxueScore
    model.ShengwuScore = s.ShengwuScore
    model.KaoshiTime = s.KaoshiTime
}

func (s *TbChengjiUpdateReq) GetId() interface{} {
	return s.Id
}

// TbChengjiGetReq 功能获取请求参数
type TbChengjiGetReq struct {
     Id int `uri:"id"`
}
func (s *TbChengjiGetReq) GetId() interface{} {
	return s.Id
}

// TbChengjiDeleteReq 功能删除请求参数
type TbChengjiDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbChengjiDeleteReq) GetId() interface{} {
	return s.Ids
}

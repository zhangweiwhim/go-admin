package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TbTermGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:tb_term" comment:"名称"`
	StartTime      string `form:"startTime"  search:"type:contains;column:start_time;table:tb_term" comment:"开始时间"`
	EndTime        string `form:"endTime"  search:"type:contains;column:end_time;table:tb_term" comment:"结束时间"`
	TbTermOrder
}

type TbTermOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:tb_term"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:tb_term"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_term"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_term"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_term"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_term"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_term"`
	StartTime string `form:"startTimeOrder"  search:"type:order;column:start_time;table:tb_term"`
	EndTime   string `form:"endTimeOrder"  search:"type:order;column:end_time;table:tb_term"`
}

func (m *TbTermGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbTermInsertReq struct {
	Id        int    `json:"-" comment:"主键编码"` // 主键编码
	Name      string `json:"name" comment:"名称"`
	StartTime string `json:"startTime" comment:"开始时间"`
	EndTime   string `json:"endTime" comment:"结束时间"`
	common.ControlBy
}

func (s *TbTermInsertReq) Generate(model *models.TbTerm) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.StartTime = s.StartTime
	model.EndTime = s.EndTime
}

func (s *TbTermInsertReq) GetId() interface{} {
	return s.Id
}

type TbTermUpdateReq struct {
	Id        int    `uri:"id" comment:"主键编码"` // 主键编码
	Name      string `json:"name" comment:"名称"`
	StartTime string `json:"startTime" comment:"开始时间"`
	EndTime   string `json:"endTime" comment:"结束时间"`
	common.ControlBy
}

func (s *TbTermUpdateReq) Generate(model *models.TbTerm) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.StartTime = s.StartTime
	model.EndTime = s.EndTime
}

func (s *TbTermUpdateReq) GetId() interface{} {
	return s.Id
}

// TbTermGetReq 功能获取请求参数
type TbTermGetReq struct {
	Id int `uri:"id"`
}

func (s *TbTermGetReq) GetId() interface{} {
	return s.Id
}

// TbTermDeleteReq 功能删除请求参数
type TbTermDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbTermDeleteReq) GetId() interface{} {
	return s.Ids
}

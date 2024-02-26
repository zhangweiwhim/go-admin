package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TbSubGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:tb_sub" comment:"名称"`
	TbSubOrder
}

type TbSubOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:tb_sub"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:tb_sub"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_sub"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_sub"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_sub"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_sub"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_sub"`
}

func (m *TbSubGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TbSubInsertReq struct {
	Id   int    `json:"-" comment:"主键编码"` // 主键编码
	Name string `json:"name" comment:"名称"`
	common.ControlBy
}

func (s *TbSubInsertReq) Generate(model *models.TbSub) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *TbSubInsertReq) GetId() interface{} {
	return s.Id
}

type TbSubUpdateReq struct {
	Id   int    `uri:"id" comment:"主键编码"` // 主键编码
	Name string `json:"name" comment:"名称"`
	common.ControlBy
}

func (s *TbSubUpdateReq) Generate(model *models.TbSub) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *TbSubUpdateReq) GetId() interface{} {
	return s.Id
}

// TbSubGetReq 功能获取请求参数
type TbSubGetReq struct {
	Id int `uri:"id"`
}

func (s *TbSubGetReq) GetId() interface{} {
	return s.Id
}

// TbSubDeleteReq 功能删除请求参数
type TbSubDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TbSubDeleteReq) GetId() interface{} {
	return s.Ids
}

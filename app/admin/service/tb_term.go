package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type TbTerm struct {
	service.Service
}

// GetPage 获取TbTerm列表
func (e *TbTerm) GetPage(c *dto.TbTermGetPageReq, p *actions.DataPermission, list *[]models.TbTerm, count *int64) error {
	var err error
	var data models.TbTerm

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TbTermService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TbTerm对象
func (e *TbTerm) Get(d *dto.TbTermGetReq, p *actions.DataPermission, model *models.TbTerm) error {
	var data models.TbTerm

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTbTerm error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TbTerm对象
func (e *TbTerm) Insert(c *dto.TbTermInsertReq) error {
	var err error
	var data models.TbTerm
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TbTermService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TbTerm对象
func (e *TbTerm) Update(c *dto.TbTermUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.TbTerm{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("TbTermService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除TbTerm
func (e *TbTerm) Remove(d *dto.TbTermDeleteReq, p *actions.DataPermission) error {
	var data models.TbTerm

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveTbTerm error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

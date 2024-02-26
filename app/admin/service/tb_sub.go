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

type TbSub struct {
	service.Service
}

// GetPage 获取TbSub列表
func (e *TbSub) GetPage(c *dto.TbSubGetPageReq, p *actions.DataPermission, list *[]models.TbSub, count *int64) error {
	var err error
	var data models.TbSub

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TbSubService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TbSub对象
func (e *TbSub) Get(d *dto.TbSubGetReq, p *actions.DataPermission, model *models.TbSub) error {
	var data models.TbSub

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTbSub error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TbSub对象
func (e *TbSub) Insert(c *dto.TbSubInsertReq) error {
	var err error
	var data models.TbSub
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TbSubService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TbSub对象
func (e *TbSub) Update(c *dto.TbSubUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.TbSub{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("TbSubService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除TbSub
func (e *TbSub) Remove(d *dto.TbSubDeleteReq, p *actions.DataPermission) error {
	var data models.TbSub

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveTbSub error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

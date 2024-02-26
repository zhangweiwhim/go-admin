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

type TbStudent struct {
	service.Service
}

// GetPage 获取TbStudent列表
func (e *TbStudent) GetPage(c *dto.TbStudentGetPageReq, p *actions.DataPermission, list *[]models.TbStudent, count *int64) error {
	var err error
	var data models.TbStudent

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TbStudentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TbStudent对象
func (e *TbStudent) Get(d *dto.TbStudentGetReq, p *actions.DataPermission, model *models.TbStudent) error {
	var data models.TbStudent

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTbStudent error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TbStudent对象
func (e *TbStudent) Insert(c *dto.TbStudentInsertReq) error {
	var err error
	var data models.TbStudent
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TbStudentService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TbStudent对象
func (e *TbStudent) Update(c *dto.TbStudentUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.TbStudent{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("TbStudentService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除TbStudent
func (e *TbStudent) Remove(d *dto.TbStudentDeleteReq, p *actions.DataPermission) error {
	var data models.TbStudent

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveTbStudent error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

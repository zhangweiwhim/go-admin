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

type TbClass struct {
	service.Service
}

// GetPage 获取TbClass列表
func (e *TbClass) GetPage(c *dto.TbClassGetPageReq, p *actions.DataPermission, list *[]models.TbClass, count *int64) error {
	var err error
	var data models.TbClass

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TbClassService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TbClass对象
func (e *TbClass) Get(d *dto.TbClassGetReq, p *actions.DataPermission, model *models.TbClass) error {
	var data models.TbClass

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTbClass error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TbClass对象
func (e *TbClass) Insert(c *dto.TbClassInsertReq) error {
    var err error
    var data models.TbClass
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TbClassService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TbClass对象
func (e *TbClass) Update(c *dto.TbClassUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.TbClass{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("TbClassService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除TbClass
func (e *TbClass) Remove(d *dto.TbClassDeleteReq, p *actions.DataPermission) error {
	var data models.TbClass

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveTbClass error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

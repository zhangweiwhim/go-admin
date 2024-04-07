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

type TbChengji struct {
	service.Service
}

// GetPage 获取TbChengji列表
func (e *TbChengji) GetPage(c *dto.TbChengjiGetPageReq, p *actions.DataPermission, list *[]models.TbChengji, count *int64) error {
	var err error
	var data models.TbChengji

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TbChengjiService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TbChengji对象
func (e *TbChengji) Get(d *dto.TbChengjiGetReq, p *actions.DataPermission, model *models.TbChengji) error {
	var data models.TbChengji

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTbChengji error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TbChengji对象
func (e *TbChengji) Insert(c *dto.TbChengjiInsertReq) error {
    var err error
    var data models.TbChengji
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TbChengjiService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TbChengji对象
func (e *TbChengji) Update(c *dto.TbChengjiUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.TbChengji{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("TbChengjiService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除TbChengji
func (e *TbChengji) Remove(d *dto.TbChengjiDeleteReq, p *actions.DataPermission) error {
	var data models.TbChengji

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveTbChengji error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

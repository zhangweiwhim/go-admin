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

type TbTeachAll struct {
	service.Service
}

// GetPage 获取TbTeachAll列表
func (e *TbTeachAll) GetPage(c *dto.TbTeachAllGetPageReq, p *actions.DataPermission, list *[]models.TbTeachAll, count *int64) error {
	var err error
	var data models.TbTeachAll

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TbTeachAllService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TbTeachAll对象
func (e *TbTeachAll) Get(d *dto.TbTeachAllGetReq, p *actions.DataPermission, model *models.TbTeachAll) error {
	var data models.TbTeachAll

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTbTeachAll error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TbTeachAll对象
func (e *TbTeachAll) Insert(c *dto.TbTeachAllInsertReq) error {
    var err error
    var data models.TbTeachAll
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TbTeachAllService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TbTeachAll对象
func (e *TbTeachAll) Update(c *dto.TbTeachAllUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.TbTeachAll{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("TbTeachAllService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除TbTeachAll
func (e *TbTeachAll) Remove(d *dto.TbTeachAllDeleteReq, p *actions.DataPermission) error {
	var data models.TbTeachAll

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveTbTeachAll error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}

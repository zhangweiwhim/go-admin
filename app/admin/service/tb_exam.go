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

type TbExam struct {
	service.Service
}

// GetPage 获取TbExam列表
func (e *TbExam) GetPage(c *dto.TbExamGetPageReq, p *actions.DataPermission, list *[]models.TbExam, count *int64) error {
	var err error
	var data models.TbExam

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TbExamService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TbExam对象
func (e *TbExam) Get(d *dto.TbExamGetReq, p *actions.DataPermission, model *models.TbExam) error {
	var data models.TbExam

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTbExam error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TbExam对象
func (e *TbExam) Insert(c *dto.TbExamInsertReq) error {
	var err error
	var data models.TbExam
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TbExamService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TbExam对象
func (e *TbExam) Update(c *dto.TbExamUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.TbExam{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("TbExamService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除TbExam
func (e *TbExam) Remove(d *dto.TbExamDeleteReq, p *actions.DataPermission) error {
	var data models.TbExam

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveTbExam error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

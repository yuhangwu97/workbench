package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"
)

type SysInterview struct {
	service.Service
}

// GetPage 获取SysInterview列表
func (e *SysInterview) GetPage(c *dto.SysInterviewPageReq, list *[]models.SysInterview, count *int64) error {
	var err error
	var data models.SysInterview

	err = e.Orm.Model(&data).
		Scopes(
			cDto.OrderDest("updated_at", false),
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s \r", err)
		return err
	}
	return nil
}

// Get 获取SysInterview对象
func (e *SysInterview) Get(d *dto.SysInterviewGetReq, model *models.SysInterview) error {
	var err error
	var data models.SysInterview

	db := e.Orm.Model(&data).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysInterview对象
func (e *SysInterview) Insert(c *dto.SysInterviewInsertReq) error {
	var err error
	var data models.SysInterview
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Update 修改SysInterview对象
func (e *SysInterview) Update(c *dto.SysInterviewUpdateReq) error {
	var err error
	var model = models.SysInterview{}
	e.Orm.First(&model, c.GetId())
	c.Generate(&model)

	db := e.Orm.Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	return nil
}

// Remove 删除SysInterview
func (e *SysInterview) Remove(d *dto.SysInterviewDeleteReq) error {
	var err error
	var data models.SysInterview

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

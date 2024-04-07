package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type TbClass struct {
	api.Api
}

// GetPage 获取班级管理列表
// @Summary 获取班级管理列表
// @Description 获取班级管理列表
// @Tags 班级管理
// @Param xueqi query string false "学期"
// @Param grade query string false "年级"
// @Param class query string false "班级"
// @Param headTeacher query string false "班主任"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TbClass}} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-class [get]
// @Security Bearer
func (e TbClass) GetPage(c *gin.Context) {
    req := dto.TbClassGetPageReq{}
    s := service.TbClass{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.TbClass, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取班级管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取班级管理
// @Summary 获取班级管理
// @Description 获取班级管理
// @Tags 班级管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.TbClass} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-class/{id} [get]
// @Security Bearer
func (e TbClass) Get(c *gin.Context) {
	req := dto.TbClassGetReq{}
	s := service.TbClass{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.TbClass

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取班级管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建班级管理
// @Summary 创建班级管理
// @Description 创建班级管理
// @Tags 班级管理
// @Accept application/json
// @Product application/json
// @Param data body dto.TbClassInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/tb-class [post]
// @Security Bearer
func (e TbClass) Insert(c *gin.Context) {
    req := dto.TbClassInsertReq{}
    s := service.TbClass{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建班级管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改班级管理
// @Summary 修改班级管理
// @Description 修改班级管理
// @Tags 班级管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.TbClassUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/tb-class/{id} [put]
// @Security Bearer
func (e TbClass) Update(c *gin.Context) {
    req := dto.TbClassUpdateReq{}
    s := service.TbClass{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改班级管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除班级管理
// @Summary 删除班级管理
// @Description 删除班级管理
// @Tags 班级管理
// @Param data body dto.TbClassDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/tb-class [delete]
// @Security Bearer
func (e TbClass) Delete(c *gin.Context) {
    s := service.TbClass{}
    req := dto.TbClassDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除班级管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}

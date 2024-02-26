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

type TbTeachAll struct {
	api.Api
}

// GetPage 获取任教管理列表
// @Summary 获取任教管理列表
// @Description 获取任教管理列表
// @Tags 任教管理
// @Param grade query string false "任教年级"
// @Param class query string false "任教班级"
// @Param subject query string false "任教学科"
// @Param techName query string false "教师名称"
// @Param techNo query string false "教师工号"
// @Param techTerm query string false "任教学期"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TbTeachAll}} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-teach-all [get]
// @Security Bearer
func (e TbTeachAll) GetPage(c *gin.Context) {
	req := dto.TbTeachAllGetPageReq{}
	s := service.TbTeachAll{}
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
	list := make([]models.TbTeachAll, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取任教管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取任教管理
// @Summary 获取任教管理
// @Description 获取任教管理
// @Tags 任教管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.TbTeachAll} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-teach-all/{id} [get]
// @Security Bearer
func (e TbTeachAll) Get(c *gin.Context) {
	req := dto.TbTeachAllGetReq{}
	s := service.TbTeachAll{}
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
	var object models.TbTeachAll

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取任教管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建任教管理
// @Summary 创建任教管理
// @Description 创建任教管理
// @Tags 任教管理
// @Accept application/json
// @Product application/json
// @Param data body dto.TbTeachAllInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/tb-teach-all [post]
// @Security Bearer
func (e TbTeachAll) Insert(c *gin.Context) {
	req := dto.TbTeachAllInsertReq{}
	s := service.TbTeachAll{}
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
		e.Error(500, err, fmt.Sprintf("创建任教管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改任教管理
// @Summary 修改任教管理
// @Description 修改任教管理
// @Tags 任教管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.TbTeachAllUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/tb-teach-all/{id} [put]
// @Security Bearer
func (e TbTeachAll) Update(c *gin.Context) {
	req := dto.TbTeachAllUpdateReq{}
	s := service.TbTeachAll{}
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
		e.Error(500, err, fmt.Sprintf("修改任教管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除任教管理
// @Summary 删除任教管理
// @Description 删除任教管理
// @Tags 任教管理
// @Param data body dto.TbTeachAllDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/tb-teach-all [delete]
// @Security Bearer
func (e TbTeachAll) Delete(c *gin.Context) {
	s := service.TbTeachAll{}
	req := dto.TbTeachAllDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除任教管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

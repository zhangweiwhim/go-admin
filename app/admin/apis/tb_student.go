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

type TbStudent struct {
	api.Api
}

// GetPage 获取学生管理列表
// @Summary 获取学生管理列表
// @Description 获取学生管理列表
// @Tags 学生管理
// @Param name query string false "名称"
// @Param grade query string false "年级"
// @Param class query string false "班级"
// @Param stuNo query string false "学号"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TbStudent}} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-student [get]
// @Security Bearer
func (e TbStudent) GetPage(c *gin.Context) {
	req := dto.TbStudentGetPageReq{}
	s := service.TbStudent{}
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
	list := make([]models.TbStudent, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取学生管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取学生管理
// @Summary 获取学生管理
// @Description 获取学生管理
// @Tags 学生管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.TbStudent} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-student/{id} [get]
// @Security Bearer
func (e TbStudent) Get(c *gin.Context) {
	req := dto.TbStudentGetReq{}
	s := service.TbStudent{}
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
	var object models.TbStudent

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取学生管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建学生管理
// @Summary 创建学生管理
// @Description 创建学生管理
// @Tags 学生管理
// @Accept application/json
// @Product application/json
// @Param data body dto.TbStudentInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/tb-student [post]
// @Security Bearer
func (e TbStudent) Insert(c *gin.Context) {
	req := dto.TbStudentInsertReq{}
	s := service.TbStudent{}
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
		e.Error(500, err, fmt.Sprintf("创建学生管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改学生管理
// @Summary 修改学生管理
// @Description 修改学生管理
// @Tags 学生管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.TbStudentUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/tb-student/{id} [put]
// @Security Bearer
func (e TbStudent) Update(c *gin.Context) {
	req := dto.TbStudentUpdateReq{}
	s := service.TbStudent{}
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
		e.Error(500, err, fmt.Sprintf("修改学生管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除学生管理
// @Summary 删除学生管理
// @Description 删除学生管理
// @Tags 学生管理
// @Param data body dto.TbStudentDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/tb-student [delete]
// @Security Bearer
func (e TbStudent) Delete(c *gin.Context) {
	s := service.TbStudent{}
	req := dto.TbStudentDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除学生管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

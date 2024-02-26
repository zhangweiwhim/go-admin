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

type TbTerm struct {
	api.Api
}

// GetPage 获取学年管理列表
// @Summary 获取学年管理列表
// @Description 获取学年管理列表
// @Tags 学年管理
// @Param name query string false "名称"
// @Param startTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TbTerm}} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-term [get]
// @Security Bearer
func (e TbTerm) GetPage(c *gin.Context) {
	req := dto.TbTermGetPageReq{}
	s := service.TbTerm{}
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
	list := make([]models.TbTerm, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取学年管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取学年管理
// @Summary 获取学年管理
// @Description 获取学年管理
// @Tags 学年管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.TbTerm} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-term/{id} [get]
// @Security Bearer
func (e TbTerm) Get(c *gin.Context) {
	req := dto.TbTermGetReq{}
	s := service.TbTerm{}
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
	var object models.TbTerm

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取学年管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建学年管理
// @Summary 创建学年管理
// @Description 创建学年管理
// @Tags 学年管理
// @Accept application/json
// @Product application/json
// @Param data body dto.TbTermInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/tb-term [post]
// @Security Bearer
func (e TbTerm) Insert(c *gin.Context) {
	req := dto.TbTermInsertReq{}
	s := service.TbTerm{}
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
		e.Error(500, err, fmt.Sprintf("创建学年管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改学年管理
// @Summary 修改学年管理
// @Description 修改学年管理
// @Tags 学年管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.TbTermUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/tb-term/{id} [put]
// @Security Bearer
func (e TbTerm) Update(c *gin.Context) {
	req := dto.TbTermUpdateReq{}
	s := service.TbTerm{}
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
		e.Error(500, err, fmt.Sprintf("修改学年管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除学年管理
// @Summary 删除学年管理
// @Description 删除学年管理
// @Tags 学年管理
// @Param data body dto.TbTermDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/tb-term [delete]
// @Security Bearer
func (e TbTerm) Delete(c *gin.Context) {
	s := service.TbTerm{}
	req := dto.TbTermDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除学年管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

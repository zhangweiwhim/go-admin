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

type TbChengji struct {
	api.Api
}

// GetPage 获取成绩管理列表
// @Summary 获取成绩管理列表
// @Description 获取成绩管理列表
// @Tags 成绩管理
// @Param name query string false "名称"
// @Param grade query string false "年级"
// @Param kaoNo query string false "考号"
// @Param kaoshiTime query string false "考试时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TbChengji}} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-chengji [get]
// @Security Bearer
func (e TbChengji) GetPage(c *gin.Context) {
    req := dto.TbChengjiGetPageReq{}
    s := service.TbChengji{}
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
	list := make([]models.TbChengji, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取成绩管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取成绩管理
// @Summary 获取成绩管理
// @Description 获取成绩管理
// @Tags 成绩管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.TbChengji} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-chengji/{id} [get]
// @Security Bearer
func (e TbChengji) Get(c *gin.Context) {
	req := dto.TbChengjiGetReq{}
	s := service.TbChengji{}
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
	var object models.TbChengji

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取成绩管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建成绩管理
// @Summary 创建成绩管理
// @Description 创建成绩管理
// @Tags 成绩管理
// @Accept application/json
// @Product application/json
// @Param data body dto.TbChengjiInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/tb-chengji [post]
// @Security Bearer
func (e TbChengji) Insert(c *gin.Context) {
    req := dto.TbChengjiInsertReq{}
    s := service.TbChengji{}
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
		e.Error(500, err, fmt.Sprintf("创建成绩管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改成绩管理
// @Summary 修改成绩管理
// @Description 修改成绩管理
// @Tags 成绩管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.TbChengjiUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/tb-chengji/{id} [put]
// @Security Bearer
func (e TbChengji) Update(c *gin.Context) {
    req := dto.TbChengjiUpdateReq{}
    s := service.TbChengji{}
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
		e.Error(500, err, fmt.Sprintf("修改成绩管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除成绩管理
// @Summary 删除成绩管理
// @Description 删除成绩管理
// @Tags 成绩管理
// @Param data body dto.TbChengjiDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/tb-chengji [delete]
// @Security Bearer
func (e TbChengji) Delete(c *gin.Context) {
    s := service.TbChengji{}
    req := dto.TbChengjiDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除成绩管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}

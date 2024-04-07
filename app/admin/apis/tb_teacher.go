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

type TbTeacher struct {
	api.Api
}

// GetPage 获取教师管理列表
// @Summary 获取教师管理列表
// @Description 获取教师管理列表
// @Tags 教师管理
// @Param name query string false "姓名"
// @Param tecNo query string false "工号"
// @Param tecNowSub query string false "任职学科"
// @Param tecNowGrade query string false "任职年级"
// @Param tecNowClass query string false "任职班级"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TbTeacher}} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-teacher [get]
// @Security Bearer
func (e TbTeacher) GetPage(c *gin.Context) {
    req := dto.TbTeacherGetPageReq{}
    s := service.TbTeacher{}
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
	list := make([]models.TbTeacher, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取教师管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取教师管理
// @Summary 获取教师管理
// @Description 获取教师管理
// @Tags 教师管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.TbTeacher} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-teacher/{id} [get]
// @Security Bearer
func (e TbTeacher) Get(c *gin.Context) {
	req := dto.TbTeacherGetReq{}
	s := service.TbTeacher{}
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
	var object models.TbTeacher

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取教师管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建教师管理
// @Summary 创建教师管理
// @Description 创建教师管理
// @Tags 教师管理
// @Accept application/json
// @Product application/json
// @Param data body dto.TbTeacherInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/tb-teacher [post]
// @Security Bearer
func (e TbTeacher) Insert(c *gin.Context) {
    req := dto.TbTeacherInsertReq{}
    s := service.TbTeacher{}
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
		e.Error(500, err, fmt.Sprintf("创建教师管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改教师管理
// @Summary 修改教师管理
// @Description 修改教师管理
// @Tags 教师管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.TbTeacherUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/tb-teacher/{id} [put]
// @Security Bearer
func (e TbTeacher) Update(c *gin.Context) {
    req := dto.TbTeacherUpdateReq{}
    s := service.TbTeacher{}
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
		e.Error(500, err, fmt.Sprintf("修改教师管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除教师管理
// @Summary 删除教师管理
// @Description 删除教师管理
// @Tags 教师管理
// @Param data body dto.TbTeacherDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/tb-teacher [delete]
// @Security Bearer
func (e TbTeacher) Delete(c *gin.Context) {
    s := service.TbTeacher{}
    req := dto.TbTeacherDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除教师管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}

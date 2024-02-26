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

type TbExam struct {
	api.Api
}

// GetPage 获取考试管理列表
// @Summary 获取考试管理列表
// @Description 获取考试管理列表
// @Tags 考试管理
// @Param examStart query string false "开始时间"
// @Param examType query string false "考试批次"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TbExam}} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-exam [get]
// @Security Bearer
func (e TbExam) GetPage(c *gin.Context) {
	req := dto.TbExamGetPageReq{}
	s := service.TbExam{}
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
	list := make([]models.TbExam, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取考试管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取考试管理
// @Summary 获取考试管理
// @Description 获取考试管理
// @Tags 考试管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.TbExam} "{"code": 200, "data": [...]}"
// @Router /api/v1/tb-exam/{id} [get]
// @Security Bearer
func (e TbExam) Get(c *gin.Context) {
	req := dto.TbExamGetReq{}
	s := service.TbExam{}
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
	var object models.TbExam

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取考试管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建考试管理
// @Summary 创建考试管理
// @Description 创建考试管理
// @Tags 考试管理
// @Accept application/json
// @Product application/json
// @Param data body dto.TbExamInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/tb-exam [post]
// @Security Bearer
func (e TbExam) Insert(c *gin.Context) {
	req := dto.TbExamInsertReq{}
	s := service.TbExam{}
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
		e.Error(500, err, fmt.Sprintf("创建考试管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改考试管理
// @Summary 修改考试管理
// @Description 修改考试管理
// @Tags 考试管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.TbExamUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/tb-exam/{id} [put]
// @Security Bearer
func (e TbExam) Update(c *gin.Context) {
	req := dto.TbExamUpdateReq{}
	s := service.TbExam{}
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
		e.Error(500, err, fmt.Sprintf("修改考试管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除考试管理
// @Summary 删除考试管理
// @Description 删除考试管理
// @Tags 考试管理
// @Param data body dto.TbExamDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/tb-exam [delete]
// @Security Bearer
func (e TbExam) Delete(c *gin.Context) {
	s := service.TbExam{}
	req := dto.TbExamDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除考试管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

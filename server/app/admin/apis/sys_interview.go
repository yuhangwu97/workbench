package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type SysInterview struct {
	api.Api
}

// GetPage
// @Summary 岗位列表数据
// @Description 获取JSON
// @Tags 岗位
// @Param postName query string false "postName"
// @Param postCode query string false "postCode"
// @Param postId query string false "postId"
// @Param status query string false "status"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [get]
// @Security Bearer
func (e SysInterview) GetPage(c *gin.Context) {
	s := service.SysInterview{}
	req := dto.SysInterviewPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.SysInterview, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get
// @Summary 获取岗位信息
// @Description 获取JSON
// @Tags 岗位
// @Param id path int true "编码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post/{postId} [get]
// @Security Bearer
func (e SysInterview) Get(c *gin.Context) {
	s := service.SysInterview{}
	req := dto.SysInterviewGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysInterview

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("岗位信息获取失败！错误详情：%s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert
// @Summary 添加岗位
// @Description 获取JSON
// @Tags 岗位
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysInterviewInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [post]
// @Security Bearer
func (e SysInterview) Insert(c *gin.Context) {
	s := service.SysInterview{}
	req := dto.SysInterviewInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("新建岗位失败！错误详情：%s", err.Error()))
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update
// @Summary 修改岗位
// @Description 获取JSON
// @Tags 岗位
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysInterviewUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post/{id} [put]
// @Security Bearer
func (e SysInterview) Update(c *gin.Context) {
	s := service.SysInterview{}
	req := dto.SysInterviewUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))

	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("岗位更新失败！错误详情：%s", err.Error()))
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除岗位
// @Description 删除数据
// @Tags 岗位
// @Param id body dto.SysInterviewDeleteReq true "请求参数"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [delete]
// @Security Bearer
func (e SysInterview) Delete(c *gin.Context) {
	s := service.SysInterview{}
	req := dto.SysInterviewDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("岗位删除失败！错误详情：%s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

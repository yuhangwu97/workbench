package dto

import (
	"go-admin/app/admin/models"
	common "go-admin/common/models"

	"go-admin/common/dto"
)

// SysInterviewPageReq 列表或者搜索使用结构体
type SysInterviewPageReq struct {
	dto.Pagination `search:"-"`
	InterviewId    int    `form:"interviewId" search:"type:exact;column:interview_id;table:sys_interview" comment:"id"`        // id
	InterviewName  string `form:"interviewName" search:"type:contains;column:interview_name;table:sys_interview" comment:"名称"` // 名称
	PostId         string `form:"postId" search:"type:contains;column:post_id;table:sys_interview" comment:"编码"`               // 编码
	Status         int    `form:"status" search:"type:exact;column:status;table:sys_interview" comment:"状态"`                   // 状态
	Remark         string `form:"remark" search:"type:exact;column:remark;table:sys_interview" comment:"备注"`                   // 备注
}

func (m *SysInterviewPageReq) GetNeedSearch() interface{} {
	return *m
}

// SysInterviewInsertReq 增使用的结构体
type SysInterviewInsertReq struct {
	InterviewId   int    `uri:"interviewId"  comment:"id"`
	InterviewName string `form:"interviewName"  comment:"名称"`
	PostId        string `form:"postId" comment:"编码"`
	Sort          int    `form:"sort" comment:"排序"`
	Status        int    `form:"status"   comment:"状态"`
	Remark        string `form:"remark"   comment:"备注"`
	common.ControlBy
}

func (s *SysInterviewInsertReq) Generate(model *models.SysInterview) {
	model.InterviewName = s.InterviewName
	model.PostId = s.PostId
	model.Status = s.Status
	model.Remark = s.Remark
	if s.ControlBy.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
	if s.ControlBy.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
}

// GetId 获取数据对应的ID
func (s *SysInterviewInsertReq) GetId() interface{} {
	return s.InterviewId
}

// SysInterviewUpdateReq 改使用的结构体
type SysInterviewUpdateReq struct {
	InterviewId   int    `uri:"id"  comment:"id"`
	InterviewName string `form:"interviewName"  comment:"名称"`
	PostId        string `form:"postId" comment:"编码"`
	Status        int    `form:"status"   comment:"状态"`
	Remark        string `form:"remark"   comment:"备注"`
	common.ControlBy
}

func (s *SysInterviewUpdateReq) Generate(model *models.SysInterview) {
	model.InterviewId = s.InterviewId
	model.InterviewName = s.InterviewName
	model.PostId = s.PostId
	model.Status = s.Status
	model.Remark = s.Remark
	if s.ControlBy.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
	if s.ControlBy.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
}

func (s *SysInterviewUpdateReq) GetId() interface{} {
	return s.InterviewId
}

// SysInterviewGetReq 获取单个的结构体
type SysInterviewGetReq struct {
	Id int `uri:"id"`
}

func (s *SysInterviewGetReq) GetId() interface{} {
	return s.Id
}

// SysInterviewDeleteReq 删除的结构体
type SysInterviewDeleteReq struct {
	Ids []int `json:"ids"`
	common.ControlBy
}

func (s *SysInterviewDeleteReq) Generate(model *models.SysInterview) {
	if s.ControlBy.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
	if s.ControlBy.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
}

func (s *SysInterviewDeleteReq) GetId() interface{} {
	return s.Ids
}

package models

type SysInterview struct {
	InterviewId   int    `gorm:"primaryKey;autoIncrement" json:"interviewId"` //候选人编号
	InterviewName string `gorm:"size:128;" json:"interviewName"`              //候选人名称
	PostId        string `gorm:"size:128;" json:"postId"`                     //岗位代码
	Status        int    `gorm:"size:8;" json:"status"`                       //状态
	Remark        string `gorm:"type:text;" json:"remark"`                    //描述
	ControlBy
	ModelTime
}

func (SysInterview) TableName() string {
	return "sys_interview"
}

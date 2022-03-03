package requestModel

import "github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"

// GetSignTaskFieldUrlReq 获取签署任务控件设置链接
type GetSignTaskFieldUrlReq struct {
	SignTaskId     string                      `json:"signTaskId,omitempty"`
	OpenUserId     string                      `json:"openUserId,omitempty"`
	UserIdentInfo  *commonModel.UserIdentInfo  `json:"userIdentInfo,omitempty"`
	UserInfoExtend *commonModel.UserInfoExtend `json:"userInfoExtend,omitempty"`
	RedirectUrl    string                      `json:"redirectUrl,omitempty"`
}

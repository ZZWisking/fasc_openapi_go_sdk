package requestModel

import (
	commonModel2 "github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"
)

// GetUserAuthUrlReq 获取个人用户授权链接
type GetUserAuthUrlReq struct {
	OpenUserId       string                      `json:"openUserId,omitempty"`
	UserIdentInfo    *commonModel2.UserIdentInfo  `json:"userIdentInfo,omitempty"`
	UserInfoExtend   *commonModel2.UserInfoExtend `json:"userInfoExtend,omitempty"`
	IdentInfoMatch   bool                        `json:"identInfoMatch"`
	RequestAuthScope []string                    `json:"requestAuthScope,omitempty"`
	RedirectUrl      string                      `json:"redirectUrl,omitempty"`
}
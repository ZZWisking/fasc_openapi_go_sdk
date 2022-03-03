package requestModel

import (
	commonModel2 "fasc_openapi_go_sdk/model/commonModel"
)

// GetCorpAuthUrlReq 获取企业用户授权链接
type GetCorpAuthUrlReq struct {
	OpenCorpId             string                       `json:"openCorpId,omitempty"`
	CorpIdentInfo          *commonModel2.CorpIdentInfo  `json:"corpIdentInfo,omitempty"`
	CorpIdentInfoMatch     bool                         `json:"corpIdentInfoMatch"`
	OperatorId             string                       `json:"operatorId,omitempty"`
	OperatorIdentInfo      *commonModel2.UserIdentInfo  `json:"operatorIdentInfo,omitempty"`
	OperatorInfoExtend     *commonModel2.UserInfoExtend `json:"operatorInfoExtend,omitempty"`
	OperatorIdentInfoMatch string                       `json:"operatorIdentInfoMatch,omitempty"`
	RequestAuthScope       []string                     `json:"requestAuthScope,omitempty"`
	RedirectUrl            string                       `json:"redirectUrl"`
}

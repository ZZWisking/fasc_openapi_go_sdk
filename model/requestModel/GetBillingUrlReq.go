package requestModel

import "github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"

// GetBillingUrlReq 获取计费链接
type GetBillingUrlReq struct {
	OpenId      *commonModel.OpenId `json:"openId,omitempty"`
	UrlType     string              `json:"urlType,omitempty"`
	RedirectUrl string              `json:"redirectUrl,omitempty"`
}

package requestModel

import "github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"

// GetDocTemplateDetailReq 获取文档模板详情
type GetDocTemplateDetailReq struct {
	OwnerId       *commonModel.OpenId `json:"ownerId,omitempty"`
	DocTemplateId string              `json:"docTemplateId,omitempty"`
}

package requestModel

import "github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"

// GetSignTemplateDetailReq 获取签署模板详情
type GetSignTemplateDetailReq struct {
	OwnerId        *commonModel.OpenId `json:"ownerId,omitempty"`
	SignTemplateId string              `json:"signTemplateId,omitempty"`
}

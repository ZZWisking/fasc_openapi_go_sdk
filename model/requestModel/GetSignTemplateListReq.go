package requestModel

import (
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"
)

// GetSignTemplateListReq 获取签署模板列表
type GetSignTemplateListReq struct {
	OwnerId      *commonModel.OpenId         `json:"ownerId,omitempty"`
	ListFilter   *SignTemplateListFilterInfo `json:"listFilter,omitempty"`
	ListPageNo   int                         `json:"listPageNo,omitempty"`
	ListPageSize int                         `json:"listPageSize,omitempty"`
}

type SignTemplateListFilterInfo struct {
	SignTemplateName string `json:"signTemplateName,omitempty"`
}

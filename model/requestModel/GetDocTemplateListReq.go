package requestModel

import "github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"

// GetDocTemplateListReq 查询文档模板列表
type GetDocTemplateListReq struct {
	OwnerId      *commonModel.OpenId   `json:"ownerId,omitempty"`
	ListFilter   *DocTemplateLisFilter `json:"listFilter,omitempty"`
	ListPageNo   int                   `json:"listPageNo,omitempty"`
	ListPageSize int                   `json:"listPageSize,omitempty"`
}

type DocTemplateLisFilter struct {
	DocTemplateName string `json:"docTemplateName,omitempty"`
}

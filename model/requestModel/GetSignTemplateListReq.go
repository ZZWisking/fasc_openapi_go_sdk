package requestModel

import (
	"fasc_openapi_go_sdk/model/commonModel"
)

type GetSignTemplateListReq struct {
	OwnerId      *commonModel.OpenId         `json:"ownerId,omitempty"`
	ListFilter   *SignTemplateListFilterInfo `json:"listFilter,omitempty"`
	ListPageNo   int                        `json:"listPageNo,omitempty"`
	ListPageSize int                        `json:"listPageSize,omitempty"`
}

type SignTemplateListFilterInfo struct {
	SignTemplateName string `json:"signTemplateName,omitempty"`
}

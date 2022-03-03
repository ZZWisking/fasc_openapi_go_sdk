package requestModel

import "github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"

// AddSignTaskDocReq 添加签署任务文档信息
type AddSignTaskDocReq struct {
	SignTaskId string                         `json:"signTaskId,omitempty"`
	Docs       *[]commonModel.SignTaskDocInfo `json:"docs,omitempty"`
}

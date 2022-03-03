package requestModel

import "fasc_openapi_go_sdk/model/commonModel"

type AddSignTaskField struct {
	DocId     int                  `json:"docId,omitempty"`
	DocFields *[]commonModel.Field `json:"docFields,omitempty"`
}

// AddSignTaskFieldReq 添加签署任务控件
type AddSignTaskFieldReq struct {
	SignTaskId string              `json:"signTaskId,omitempty"`
	Fields     *[]AddSignTaskField `json:"fields,omitempty"`
}

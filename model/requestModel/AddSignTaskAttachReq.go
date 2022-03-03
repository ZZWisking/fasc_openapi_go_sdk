package requestModel

import commonModel2 "fasc_openapi_go_sdk/model/commonModel"

// AddSignTaskAttachReq 添加签署任务附件

type AddSignTaskAttachReq struct {
	SignTaskId string                          `json:"signTaskId,omitempty"`
	Attachs    *[]commonModel2.SignTaskAttachs `json:"attachs,omitempty"`
}

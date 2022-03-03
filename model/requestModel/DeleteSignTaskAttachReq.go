package requestModel

// DeleteSignTaskAttachReq 移除签署任务附件
type DeleteSignTaskAttachReq struct {
	SignTaskId string `json:"signTaskId,omitempty"`
	AttachIds  []int  `json:"attachIds,omitempty"`
}

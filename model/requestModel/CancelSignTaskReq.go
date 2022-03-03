package requestModel

// CancelSignTaskReq 撤销签署任务
type CancelSignTaskReq struct {
	SignTaskId string `json:"signTaskId,omitempty"`
}

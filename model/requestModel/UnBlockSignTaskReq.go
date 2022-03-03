package requestModel

// UnBlockSignTaskReq 解除阻塞签署任务
type UnBlockSignTaskReq struct {
	SignTaskId string `json:"signTaskId,omitempty"`
	ActorType  string `json:"actorType,omitempty"`
	ActorId    string `json:"actorId,omitempty"`
}

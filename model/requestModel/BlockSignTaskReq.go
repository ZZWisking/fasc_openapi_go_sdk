package requestModel

// BlockSignTaskReq 阻塞签署任务
type BlockSignTaskReq struct {
	SignTaskId string `json:"signTaskId,omitempty"`
	ActorType  string `json:"actorType,omitempty"`
	ActorId    string `json:"actorId,omitempty"`
}

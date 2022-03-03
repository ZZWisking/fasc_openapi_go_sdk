package requestModel

// DeleteSignTaskActorReq 移除任务参与方
type DeleteSignTaskActorReq struct {
	SignTaskId string   `json:"signTaskId,omitempty"`
	ActorType  string   `json:"actorType,omitempty"`
	ActorIds   []string `json:"actorIds,omitempty"`
}

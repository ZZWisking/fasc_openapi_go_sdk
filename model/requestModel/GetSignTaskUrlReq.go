package requestModel

// GetSignTaskUrlReq 获取签署任务链接
type GetSignTaskUrlReq struct {
	SignTaskId  string `json:"signTaskId,omitempty"`
	ActorType   string `json:"actorType,omitempty"`
	ActorId     string `json:"actorId,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
}

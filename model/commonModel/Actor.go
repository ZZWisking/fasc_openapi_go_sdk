package commonModel

// Actor 签署任务参与方
type Actor struct {
	ActorType      string        `json:"actorType,omitempty"`
	ActorId        string        `json:"actorId,omitempty"`
	ActorIdentType string        `json:"actorIdentType,omitempty"`
	ActorUser      *ActorUser    `json:"actorUser,omitempty"`
	ActorCorp      *ActorCorp    `json:"actorCorp,omitempty"`
	Notification   *Notification `json:"notification,omitempty"`
}

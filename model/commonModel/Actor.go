package commonModel


type Actor struct {
	ActorType string  `json:"actorType,omitempty"`
	ActorId string  `json:"actorId,omitempty"`
	ActorIdentType string       `json:"actorIdentType,omitempty"`
	ActorUser      *ActorUser    `json:"actorUser,omitempty"`
	ActorCorp      *ActorCorp    `json:"actorCorp,omitempty"`
	Notification   *Notification `json:"notification,omitempty"`
}



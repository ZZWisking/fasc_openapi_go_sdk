package commonModel

// FillActorInfo 填充方列表
type FillActorInfo struct {
	FillActor   *Actor            `json:"fillActor,omitempty"`
	ActorFields *[]FieldValueInfo `json:"actorFields,omitempty"`
}

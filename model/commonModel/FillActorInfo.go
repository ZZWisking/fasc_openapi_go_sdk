package commonModel

type FillActorInfo struct{
	FillActor   *Actor          `json:"fillActor,omitempty"`
	ActorFields *[]FieldValueInfo `json:"actorFields,omitempty"`
}

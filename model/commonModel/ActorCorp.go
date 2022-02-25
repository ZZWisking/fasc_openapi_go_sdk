package commonModel

type ActorCorp struct {
	ActorCorpId        string         `json:"actorCorpId,omitempty"`
	CorpIdentInfo      *CorpIdentInfo  `json:"corpIdentInfo,omitempty"`
	CorpInfoExtend     *CorpInfoExtend `json:"corpInfoExtend,omitempty"`
	OperatorId         string         `json:"operatorId,omitempty"`
	OperatorIdentInfo  *UserIdentInfo  `json:"operatorIdentInfo,omitempty"`
	OperatorInfoExtend *UserInfoExtend `json:"operatorInfoExtend,omitempty"`
}

package commonModel

type ActorUser struct {
	ActorUserId    string         `json:"actorUserId,omitempty"`
	UserIdentInfo  *UserIdentInfo  `json:"userIdentInfo,omitempty"`
	UserInfoExtend *UserInfoExtend `json:"userInfoExtend,omitempty"`
}

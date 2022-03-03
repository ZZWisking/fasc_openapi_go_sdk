package commonModel

// ActorUser 个人用户参与方
type ActorUser struct {
	ActorUserId    string          `json:"actorUserId,omitempty"`
	UserIdentInfo  *UserIdentInfo  `json:"userIdentInfo,omitempty"`
	UserInfoExtend *UserInfoExtend `json:"userInfoExtend,omitempty"`
}

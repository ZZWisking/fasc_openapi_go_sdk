package requestModel

// UserUpdateReq 更新个人用户
type UserUpdateReq struct {
	OpenUserId     string `json:"openUserId,omitempty"`
	ClientUserId   string `json:"clientUserId,omitempty"`
	ClientUserName string `json:"clientUserName,omitempty"`
}

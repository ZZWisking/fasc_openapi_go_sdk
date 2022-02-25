package requestModel

// UserAddReq 添加个人用户
type UserAddReq struct {
	ClientUserId string `json:"clientUserId,omitempty"`
	ClientUserName string `json:"clientUserName,omitempty"`
}


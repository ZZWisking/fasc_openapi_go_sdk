package commonModel


type UserIdentInfo struct {
	UserName string    `json:"userName,omitempty"`
	IdentType string    `json:"identType,omitempty"`
	IdentNo string    `json:"identNo,omitempty"`
}
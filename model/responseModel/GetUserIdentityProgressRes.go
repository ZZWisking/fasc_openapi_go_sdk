package responseModel

type GetUserIdentityProgressRes struct {
	RequestId string                         `json:"requestId"`
	Code      string                         `json:"code"`
	Msg       string                         `json:"msg"`
	Data      GetUserIdentityProgressResData `json:"data"`
}

type GetUserIdentityProgressResData struct {
	IdentProcessStatus string `json:"identProcessStatus"`
	IdentUpdateTime    string `json:"identUpdateTime"`
	IdentMethod        string `json:"identMethod"`
	IdentFailedReason  string `json:"identFailedReason"`
}

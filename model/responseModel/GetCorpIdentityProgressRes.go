package responseModel

type GetCorpIdentityProgressRes struct {
	RequestId string                         `json:"requestId"`
	Code      string                         `json:"code"`
	Msg       string                         `json:"msg"`
	Data      GetCorpIdentityProgressResData `json:"data"`
}

type GetCorpIdentityProgressResData struct {
	IdentProgressInfos []CorpIdentProgressInfo `json:"identProgressInfos"`
}

type CorpIdentProgressInfo struct {
	CorpIdentProcessStatus     string `json:"corpIdentProcessStatus"`
	CorpIdentUpdateTime        string `json:"corpIdentUpdateTime"`
	CorpIdentMethod            string `json:"corpIdentMethod"`
	CorpIdentFailedReason      string `json:"corpIdentFailedReason"`
	OperatorId                 string `json:"operatorId"`
	OperatorIdentProcessStatus string `json:"operatorIdentProcessStatus"`
	OperatorIdentMethod        string `json:"operatorIdentMethod"`
	OperatorIdentFailedReason  string `json:"operatorIdentFailedReason"`
}

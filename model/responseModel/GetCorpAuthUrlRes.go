package responseModel

type AuthUrlResData struct {
	EUrl string `json:"eUrl"`
}

type GetCorpAuthUrlRes struct {
	RequestId string         `json:"requestId"`
	Code      string         `json:"code"`
	Msg       string         `json:"msg"`
	Data      AuthUrlResData `json:"data"`
}

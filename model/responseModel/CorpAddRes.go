package responseModel

type CorpAddRes struct {
	RequestId string  `json:"requestId"`
	Code string  `json:"code"`
	Msg  string         `json:"msg"`
	Data CorpAddResData `json:"data"`
}
type CorpAddResData struct {
	OpenCorpId string `json:"openCorpId"`
}

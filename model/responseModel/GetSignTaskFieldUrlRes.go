package responseModel

type SignTaskFieldUrlData struct {
	EUrl     string `json:"eUrl"`
	CloudUrl string `json:"cloudUrl"`
}

type GetSignTaskFieldUrlRes struct {
	RequestId string               `json:"requestId"`
	Code      string               `json:"code"`
	Msg       string               `json:"msg"`
	Data      SignTaskFieldUrlData `json:"data"`
}

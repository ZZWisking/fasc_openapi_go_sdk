package responseModel

type SignTaskUrlResData struct {
	EUrl     string `json:"eUrl"`
	CloudUrl string `json:"cloudUrl"`
}

type GetSignTaskUrlRes struct {
	RequestId string             `json:"requestId"`
	Code      string             `json:"code"`
	Msg       string             `json:"msg"`
	Data      SignTaskUrlResData `json:"data"`
}

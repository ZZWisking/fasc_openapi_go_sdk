package responseModel

type SignTaskAddResData struct {
	SignTaskId string  `json:"signTaskId"`
}

type CreateSignTaskWithTemplateRes struct{
	RequestId string  `json:"requestId"`
	Code string  `json:"code"`
	Msg  string             `json:"msg"`
	Data SignTaskAddResData `json:"data"`
}




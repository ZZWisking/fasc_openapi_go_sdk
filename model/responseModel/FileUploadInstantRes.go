package responseModel

type FileUploadInstantRes struct {
	RequestId string                   `json:"requestId"`
	Code      string                   `json:"code"`
	Msg       string                   `json:"msg"`
	Data      FileUploadInstantResData `json:"data"`
}

type FileUploadInstantResData struct {
	FileId string `json:"fileId"`
}

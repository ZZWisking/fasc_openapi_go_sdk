package responseModel

type GetCorpManageUrlRes struct {
	RequestId string              `json:"requestId"`
	Code      string              `json:"code"`
	Msg       string              `json:"msg"`
	Data      GetManageUrlResData `json:"data"`
}

type GetManageUrlResData struct {
	EUrl string `json:"eUrl"`
}

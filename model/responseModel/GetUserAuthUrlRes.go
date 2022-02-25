package responseModel


type GetUserAuthUrlRes struct {
	RequestId string  `json:"requestId"`
	Code string  `json:"code"`
	Msg  string             `json:"msg"`
	Data UserAuthUrlResData `json:"data"`
}
type UserAuthUrlResData struct {
	EUrl string  `json:"eUrl"`
}

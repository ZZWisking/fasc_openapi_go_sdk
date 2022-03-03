package responseModel

type UserAddRes struct {
	RequestId string         `json:"requestId"`
	Code      string         `json:"code"`
	Msg       string         `json:"msg"`
	Data      UserAddResData `json:"data"`
}

type UserAddResData struct {
	OpenUserId string `json:"openUserId"`
}

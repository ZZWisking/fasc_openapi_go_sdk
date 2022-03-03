package responseModel

import "github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"

type GetUserIdentifyInfoRes struct {
	RequestId string                     `json:"requestId"`
	Code      string                     `json:"code"`
	Msg       string                     `json:"msg"`
	Data      GetUserIdentifyInfoResData `json:"data"`
}

type GetUserIdentifyInfoResData struct {
	OpenUserId          string                     `json:"openUserId"`
	IdentStatus         string                     `json:"identStatus"`
	UserIdentInfo       commonModel.UserIdentInfo  `json:"userIdentInfo"`
	UserIdentInfoExtend commonModel.UserInfoExtend `json:"userIdentInfoExtend"`
	identMethod         string                     `json:"identMethod"`
	identSubmitTime     string                     `json:"identSubmitTime"`
	identSuccessTime    string                     `json:"identSuccessTime"`
}

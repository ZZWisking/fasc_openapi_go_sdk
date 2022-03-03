package responseModel

import "fasc_openapi_go_sdk/model/commonModel"

type GetCorpIdentityInfoRes struct {
	RequestId string                     `json:"requestId"`
	Code      string                     `json:"code"`
	Msg       string                     `json:"msg"`
	Data      GetCorpIdentityInfoResData `json:"data"`
}

type GetCorpIdentityInfoResData struct {
	OpenCorpId          string                     `json:"openCorpId"`
	CorpIdentStatus     string                     `json:"corpIdentStatus"`
	CorpIdentInfo       commonModel.CorpIdentInfo  `json:"corpIdentInfo"`
	CorpIdentInfoExtend commonModel.CorpInfoExtend `json:"corpIdentInfoExtend"`
	CorpIdentMethod     string                     `json:"corpIdentMethod"`
	OperatorType        string                     `json:"operatorType"`
	OperatorId          string                     `json:"operatorId"`
	OperatorIdentMethod string                     `json:"operatorIdentMethod"`
	IdentSubmitTime     string                     `json:"identSubmitTime"`
	IdentSuccessTime    string                     `json:"identSuccessTime"`
}

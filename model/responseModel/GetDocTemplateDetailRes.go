package responseModel

import "github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"

type GetDocTemplateDetailRes struct {
	RequestId string                      `json:"requestId"`
	Code      string                      `json:"code"`
	Msg       string                      `json:"msg"`
	Data      GetDocTemplateDetailResData `json:"data"`
}

type GetDocTemplateDetailResData struct {
	DocTemplateId     string            `json:"docTemplateId"`
	DocTemplateName   string            `json:"docTemplateName"`
	DocTemplateStatus string            `json:"docTemplateStatus"`
	DocFields         []commonModel.Field `json:"docFields"`
}

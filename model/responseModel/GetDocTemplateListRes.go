package responseModel

type GetDocTemplateListRes struct {
	RequestId string                    `json:"requestId"`
	Code      string                    `json:"code"`
	Msg       string                    `json:"msg"`
	Data      GetDocTemplateListResData `json:"data"`
}

type GetDocTemplateListResData struct {
	DocTemplates []DocTemplates `json:"docTemplates"`
}

type DocTemplates struct {
	DocTemplateId     string `json:"docTemplateId"`
	DocTemplateName   string `json:"docTemplateName"`
	DocTemplateStatus string `json:"docTemplateStatus"`
	CreateTime        string `json:"createTime"`
	UpdateTime        string `json:"updateTime"`
	ListPageNo        int    `json:"listPageNo"`
	CountInPage       int    `json:"countInPage"`
	ListPageCount     int    `json:"listPageCount"`
	TotalCount        int    `json:"totalCount"`
}

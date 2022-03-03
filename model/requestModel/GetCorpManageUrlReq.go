package requestModel

// GetCorpManageUrlReq 获取企业管理链接
type GetCorpManageUrlReq struct {
	OpenCorpId  string `json:"openCorpId,omitempty"`
	OperatorId  string `json:"operatorId,omitempty"`
	ResourceId  string `json:"resourceId,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
}

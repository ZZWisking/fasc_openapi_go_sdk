package requestModel

// GetPageManageUrlReq 获取模板管理链接
type GetPageManageUrlReq struct {
	OpenCorpId  string `json:"openCorpId,omitempty"`
	RedirectUrl string `json:"redirectUrl,omitempty"`
}

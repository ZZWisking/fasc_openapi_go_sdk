package requestModel

// CorpUpdateReq 更新企业用户
type CorpUpdateReq struct {
	OpenCorpId     string `json:"openCorpId,omitempty"`
	ClientCorpId   string `json:"clientCorpId,omitempty"`
	ClientCorpName string `json:"clientCorpName,omitempty"`
}

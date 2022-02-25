package requestModel

// CorpAddReq 添加企业用户
type CorpAddReq struct {
	ClientCorpId string   `json:"clientCorpId,omitempty"`
	ClientCorpName string   `json:"clientCorpName,omitempty"`
}

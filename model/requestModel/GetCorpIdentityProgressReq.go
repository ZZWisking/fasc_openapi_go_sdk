package requestModel

// GetCorpIdentityProgressReq 查询企业用户实名认证进度
type GetCorpIdentityProgressReq struct {
	OpenCorpId string `json:"openCorpId,omitempty"`
	OperatorId string `json:"operatorId,omitempty"`
}

package commonModel

// OpenId 统一标识应用系统上的用户(个人或企业)。
type OpenId struct {
	IdType     string `json:"idType,omitempty"`
	OpenCorpId string `json:"openCorpId,omitempty"`
	OpenUserId string `json:"openUserId,omitempty"`
}

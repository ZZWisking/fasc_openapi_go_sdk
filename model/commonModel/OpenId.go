package commonModel

type OpenId struct {
	IdType string  `json:"idType,omitempty"`
	OpenCorpId string `json:"openCorpId,omitempty"`
	OpenUserId string `json:"openUserId,omitempty"`
}

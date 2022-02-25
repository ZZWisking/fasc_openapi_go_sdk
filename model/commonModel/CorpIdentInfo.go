package commonModel


type CorpIdentInfo struct {
	CorpName string   `json:"corpName,omitempty"`
	CorpIdentType string   `json:"corpIdentType,omitempty"`
	CorpIdentNo string   `json:"corpIdentNo,omitempty"`
	LegalRepName string  `json:"legalRepName,omitempty"`
}
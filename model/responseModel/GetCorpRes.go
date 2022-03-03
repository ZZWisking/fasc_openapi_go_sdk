package responseModel

type GetCorpRes struct {
	RequestId string         `json:"requestId"`
	Code      string         `json:"code"`
	Msg       string         `json:"msg"`
	Data      GetCorpResData `json:"data"`
}

type GetCorpResData struct {
	ClientCorpId    string   `json:"clientCorpId"`
	ClientCorpName  string   `json:"clientCorpName"`
	BindingStatus   string   `json:"bindingStatus"`
	AuthScope       []string `json:"authScope"`
	IdentStatus     string   `json:"identStatus"`
	AvailableStatus string   `json:"availableStatus"`
}

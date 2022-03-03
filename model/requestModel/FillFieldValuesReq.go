package requestModel

type DocFieldValue struct {
	DocId      int    `json:"docId"`
	FieldId    string `json:"fieldId,omitempty"`
	FieldValue string `json:"fieldValue,omitempty"`
}

// FillFieldValuesReq 填写签署任务控件内容
type FillFieldValuesReq struct {
	SignTaskId     string           `json:"signTaskId,omitempty"`
	DocFieldValues *[]DocFieldValue `json:"docFieldValues,omitempty"`
}

package commonModel

// SignTaskDocInfo 签署任务文档信息
type SignTaskDocInfo struct {
	DocId         int      `json:"docId,omitempty"`
	DocName       string   `json:"docName,omitempty"`
	DocFileId     string   `json:"docFileId,omitempty"`
	DocTemplateId string   `json:"docTemplateId,omitempty"`
	DocFields     *[]Field `json:"docFields,omitempty"`
}

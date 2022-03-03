package commonModel

// SignTaskAttachs 签署任务附件

type SignTaskAttachs struct {
	AttachId     int    `json:"attachId,omitempty"`
	AttachName   string `json:"attachName,omitempty"`
	AttachFileId string `json:"attachFileId,omitempty"`
}

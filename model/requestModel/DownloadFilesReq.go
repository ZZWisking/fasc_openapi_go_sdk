package requestModel

// DownloadFilesReq 下载签署任务文档
type  DownloadFilesReq struct {
	SignTaskId string  `json:"signTaskId,omitempty"`
	FileType  string  `json:"fileType,omitempty"`
	Id int  `json:"id,omitempty"`
}

package requestModel

import "github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"

// FileUploadInstantReq 上传即时文件
type FileUploadInstantReq struct {
	OwnerId         *commonModel.OpenId `json:"ownerId,omitempty"`
	FileType        string              `json:"fileType,omitempty"`
	FileName        string              `json:"fileName,omitempty"`
	FileUrl         string              `json:"fileUrl,omitempty"`
	FileContentHash string              `json:"fileContentHash,omitempty"`
}

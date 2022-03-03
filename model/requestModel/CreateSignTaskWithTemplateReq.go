package requestModel

import (
	commonModel2 "fasc_openapi_go_sdk/model/commonModel"
)

// CreateSignTaskWithTemplateReq 创建签署任务基于签署模板
type CreateSignTaskWithTemplateReq struct {
	SignTaskSubject  string                        `json:"signTaskSubject,omitempty"`
	Initiator        *commonModel2.OpenId          `json:"initiator,omitempty"`
	ExpiresTime      string                        `json:"expiresTime,omitempty"`
	AutoInitiate     bool                          `json:"autoInitiate"`
	AutoFillFinalize bool                          `json:"autoFillFinalize"`
	AutoFinish       bool                          `json:"autoFinish"`
	BusinessScene    *BusinessSceneInfo            `json:"businessScene,omitempty"`
	SignTemplateId   string                        `json:"signTemplateId,omitempty"`
	FillActors       *[]commonModel2.FillActorInfo `json:"fillActors,omitempty"`
	SignActors       *[]commonModel2.SignActorInfo `json:"signActors,omitempty"`
	CcActors         *[]commonModel2.CcActorInfo   `json:"ccActors,omitempty"`
}

type BusinessSceneInfo struct {
	BusinessId       string `json:"businessId,omitempty"`
	TransReferenceId string `json:"transReferenceId,omitempty"`
}

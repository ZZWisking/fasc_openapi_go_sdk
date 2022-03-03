package requestModel

import commonModel2 "github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"

// CreateSignTaskReq 创建签署任务
type CreateSignTaskReq struct {
	SignTaskSubject  string                          `json:"signTaskSubject,omitempty"`
	Initiator        *commonModel2.OpenId            `json:"initiator,omitempty"`
	ExpiresTime      string                          `json:"expiresTime,omitempty"`
	AutoInitiate     bool                            `json:"autoInitiate"`
	AutoFillFinalize bool                            `json:"autoFillFinalize"`
	AutoFinish       bool                            `json:"autoFinish"`
	FillInOrder      bool                            `json:"fillInOrder"`
	SignInOrder      bool                            `json:"signInOrder"`
	BusinessScene    *BusinessSceneInfo              `json:"businessScene,omitempty"`
	Docs             *[]commonModel2.SignTaskDocInfo `json:"docs,omitempty"`
	Attachs          *[]commonModel2.SignTaskAttachs `json:"attachs,omitempty"`
	FillActors       *[]commonModel2.FillActorInfo   `json:"fillActors,omitempty"`
	SignActors       *[]commonModel2.SignActorInfo   `json:"signActors,omitempty"`
	CcActors         *[]commonModel2.CcActorInfo     `json:"ccActors,omitempty"`
}

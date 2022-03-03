package requestModel

import commonModel2 "fasc_openapi_go_sdk/model/commonModel"

type SignTaskFillActor struct {
	FillActor   *commonModel2.Actor            `json:"fillActor,omitempty"`
	OrderNo     int                            `json:"orderNo,omitempty"`
	ActorFields *[]commonModel2.FieldValueInfo `json:"actorFields,omitempty"`
}

type SignTaskSignActor struct {
	SignActor         *commonModel2.Actor            `json:"signActor,omitempty"`
	OrderNo           int                            `json:"orderNo,omitempty"`
	BlockHere         bool                           `json:"blockHere"`
	RequestVerifyFree bool                           `json:"requestVerifyFree"`
	VerifyMethods     []string                       `json:"verifyMethods,omitempty"`
	CorpOperatorSign  bool                           `json:"corpOperatorSign"`
	SignerSignMethod  string                         `json:"signerSignMethod,omitempty"`
	ActorFields       *[]commonModel2.FieldValueInfo `json:"actorFields,omitempty"`
}

// AddSignTaskActorReq 添加签署任务参与方

type AddSignTaskActorReq struct {
	SignTaskId string                      `json:"signTaskId,omitempty"`
	FillActors *[]SignTaskFillActor        `json:"fillActors,omitempty"`
	SignActors *[]SignTaskSignActor        `json:"signActors,omitempty"`
	CcActors   *[]commonModel2.CcActorInfo `json:"ccActors,omitempty"`
}

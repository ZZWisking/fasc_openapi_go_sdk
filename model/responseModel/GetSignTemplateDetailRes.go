package responseModel

import "fasc_openapi_go_sdk/model/commonModel"

type GetSignTemplateDetailReq struct {
	RequestId string                       `json:"requestId"`
	Code      string                       `json:"code"`
	Msg       string                       `json:"msg"`
	Data      GetSignTemplateDetailReqData `json:"data"`
}

type GetSignTemplateDetailReqData struct {
	SignTemplateId     string      `json:"signTemplateId"`
	SignTemplateName   string      `json:"signTemplateName"`
	SignTemplateStatus string      `json:"signTemplateStatus"`
	Docs               []Docs      `json:"docs"`
	FillInOrder        bool        `json:"fillInOrder"`
	SignInOrder        bool        `json:"signInOrder"`
	FillActors         []FillActor `json:"fillActors"`
	SignActors         []SignActor `json:"signActors"`
}

type Docs struct {
	DocId     int                 `json:"docId"`
	DocName   string              `json:"docName"`
	DocFields []commonModel.Field `json:"docFields"`
}

type FillActor struct {
	ActorId         string       `json:"actorId"`
	OrderNo         int          `json:"orderNo"`
	ActorIdentType  string       `json:"actorIdentType"`
	FillActorFields []ActorField `json:"fillActorFields"`
}

type ActorField struct {
	FieldDocId int    `json:"fieldDocId"`
	FieldId    string `json:"fieldId"`
}

type SignActor struct {
	ActorId          string `json:"actorId"`
	OrderNo          int    `json:"orderNo"`
	ActorIdentType   string `json:"actorIdentType"`
	CorpOperatorSign bool   `json:"corpOperatorSign"`
	SignerSignMethod string `json:"signerSignMethod"`
	SignActorFields  []ActorField
}

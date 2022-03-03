package responseModel

import (
	commonModel2 "github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"
)

type SignTaskDetailDocInfo struct {
	DocId   int    `json:"docId"`
	DocName string `json:"docName"`
}

type SignTaskDetailAttachs struct {
	AttachId   int    `json:"attachId"`
	AttachName string `json:"attachName"`
}

type SignTaskDetailFillActors struct {
	FillActor       commonModel2.Actor `json:"fillActor"`
	OrderNo         int                `json:"orderNo"`
	FillActorStatus string             `json:"fillActorStatus"`
	ActionTime      string             `json:"actionTime"`
}

type SignTaskDetailSignActors struct {
	SignActor       commonModel2.Actor `json:"signActor"`
	OrderNo         int                `json:"orderNo"`
	SignActorStatus string             `json:"signActorStatus"`
	ActionTime      string             `json:"actionTime"`
}

type SignTaskDetailResData struct {
	SignTaskId      string                     `json:"signTaskId"`
	SignTaskSubject string                     `json:"signTaskSubject"`
	SignTaskStatus  string                     `json:"signTaskStatus"`
	Docs            []SignTaskDetailDocInfo    `json:"docs"`
	Attachs         []SignTaskDetailAttachs    `json:"attachs"`
	FillActors      []SignTaskDetailFillActors `json:"fillActors"`
	SignActors      []SignTaskDetailSignActors `json:"signActors"`
	CcActors        []commonModel2.CcActorInfo `json:"ccActors"`
}

type GetSignTaskDetailRes struct {
	RequestId string                `json:"requestId"`
	Code      string                `json:"code"`
	Msg       string                `json:"msg"`
	Data      SignTaskDetailResData `json:"data"`
}

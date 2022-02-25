package client

import (
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/commonModel"
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/requestModel"
	"fmt"
	"os"
	"strings"
	"testing"
)

const(
	ContentTypeJson ="application/json"
	ContentTypeZip = "application/zip"
	ContentTypeFile = "application/*"
	AppId = ""
	AppSecret = ""
	ServerUrl = ""
)

//fasc_openapi_sdk 单元测试
func TestOpenApiClient_GetAuthToken(t *testing.T) {
	//获取服务访问凭证
	 c := NewClient(AppId,AppSecret,ServerUrl)
   	 accessTokenRes := c.GetAuthToken()

	 //添加个人用户
	  userAddReq := &requestModel.UserAddReq{ClientUserId: "1111111",ClientUserName: "zhangzhiwei"}
	  userAddResult := c.GetUserAddResponse(userAddReq,accessTokenRes.Data.AccessToken)
	  //tokenReuslt := gjson.Get(userAddResult, "data.openUserId")
	  fmt.Println(userAddResult)

	 //获取个人用户授权链接
	 getUserAuthUrlReq := &requestModel.GetUserAuthUrlReq{OpenUserId:"d8080557bc7a40e99d13cb1150bb0b44" }
	 getUserAuthUrlRes := c.GetUserAuthUrlResponse(getUserAuthUrlReq,accessTokenRes.Data.AccessToken)
	 fmt.Println(getUserAuthUrlRes)

	 //添加企业用户
	 corpAddReq := &requestModel.CorpAddReq{ClientCorpId: "1314520",ClientCorpName: "永诚保险"}
	 corpAddRes := c.GetCorpAddResponse(corpAddReq,accessTokenRes.Data.AccessToken)
	 fmt.Println(corpAddRes)

	 //获取企业用户授权链接
	 getCorpAuthUrlReq := &requestModel.GetCorpAuthUrlReq{OpenCorpId:"a0b066225997491e8c4a28bf3e2ed5ac",OperatorId: "8b6e64b9a0b54f7787c95b4e9230a323"}
	 getCorpAuthUrlRes := c.GetCorpAuthUrlResponse(getCorpAuthUrlReq,accessTokenRes.Data.AccessToken)
	 fmt.Println(getCorpAuthUrlRes)

	 //获取签署模拟列表
	 //1.(如果需要参数)示例
	 //openId := &commonModel.OpenId{IdType: "person",OpenUserId:"8a3cb72ef85f4f62a6e82935eb41908c" }
	 //listFilter := &requestModel.SignTemplateListFilterInfo{SignTemplateName: "hangman"}
	 //getSignTemplateListReq := &requestModel.GetSignTemplateListReq{OwnerId:openId,ListFilter: listFilter,ListPageNo: 1,ListPageSize: 1}
	 //2.(不需要参数)示例
	 getSignTemplateListReq := &requestModel.GetSignTemplateListReq{ListPageNo: 1,ListPageSize: 1}
	 getSignTemplateListRes := c.GetSignTemplateListResponse(getSignTemplateListReq,accessTokenRes.Data.AccessToken)
	 fmt.Println(getSignTemplateListRes)

	 //创建签署任务基于签署模板
	 initiator := &commonModel.OpenId{IdType: "person",OpenUserId:"8a3cb72ef85f4f62a6e82935eb41908c" }
	// businessScene := &requestModel.BusinessSceneInfo{BusinessId: ""} //(需要时再添加)

	 notification := &commonModel.Notification{ SendNotification: false}
	 fillActor :=  &commonModel.Actor{ActorType: "filler",ActorId: "111",ActorIdentType: "person",Notification: notification}

	 actorField1 := commonModel.FieldValueInfo{FieldDocId: 111,FieldValue: ""}
	 actorField2 := commonModel.FieldValueInfo{FieldDocId: 222,FieldValue: ""}
	 actorFieldsArr := &[]commonModel.FieldValueInfo{actorField1,actorField2}

	 fillActor1 := commonModel.FillActorInfo{FillActor: fillActor,ActorFields: actorFieldsArr}
	 fillActorArr := &[]commonModel.FillActorInfo{fillActor1}

	 createSignTaskWithTemReq := &requestModel.CreateSignTaskWithTemplateReq{SignTaskSubject: "基于golang创建签署任务",SignTemplateId: "1635243428580177021",Initiator:initiator,FillActors: fillActorArr}
	 createSignTaskWithTemRes := c.CreateSignTaskWithTemplate(createSignTaskWithTemReq,accessTokenRes.Data.AccessToken)
	 fmt.Println(createSignTaskWithTemRes)

	 //获取签署任务详情
	 signTaskDetailReq := &requestModel.GetSignTaskDetailReq{SignTaskId: "1644551524323139545"}
	 signTaskDetailRes := c.GetSignTaskDetailResponse(signTaskDetailReq,accessTokenRes.Data.AccessToken)
	 fmt.Println(signTaskDetailRes)

	 //获取签署任务链接
	 signTaskUrlReq := &requestModel.GetSignTaskUrlReq{SignTaskId: "1644551524323139545",ActorType: "initiator"}
	 signTaskUrlRes := c.GetSignTaskUrlResponse(signTaskUrlReq,accessTokenRes.Data.AccessToken)
	 fmt.Println(signTaskUrlRes)

	 //下载签署任务文档
	 downloadFilesReq := &requestModel.DownloadFilesReq{SignTaskId: "1640935237063137559"}
	 downloadFilesRes := c.DownloadSignTaskFiles(downloadFilesReq,accessTokenRes.Data.AccessToken)

	 if find := strings.Contains(downloadFilesRes.ContentType, ContentTypeJson); find{
		 fmt.Println(downloadFilesRes.Data)
	 }
	 if find := strings.Contains(downloadFilesRes.ContentType, ContentTypeZip); find{
		 file,err := os.Create("./document.zip")
		 if err != nil{
		 }
		 _,err = file.Write(downloadFilesRes.Content)
		 if err != nil{
			 fmt.Println("文件写入失败")
		 }
	 }
	 if find := strings.Contains(downloadFilesRes.ContentType,ContentTypeFile);find {
	 	 //此处请创建自己上传的文件类型文件
	 }
}

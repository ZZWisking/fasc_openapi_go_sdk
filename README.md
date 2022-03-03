版本:go 1.16

调用示例: GetToken_test.go文件

包导入示例：（包地址为http://github.com官当仓库地址）

import (
"fmt"
"github.com/ZZWisking/github.com/ZZWisking/fasc_openapi_go_sdk/client"
)

const(
ContentTypeJson ="application/json"
ContentTypeZip = "application/zip"
ContentTypeFile = "application/*"
AppId = ""
AppSecret = ""
ServerUrl = ""
)

//开始调用版本sdk //1.获取token c := client.NewClient("","","")
//获取服务访问凭证 c := NewClient(AppId, AppSecret, ServerUrl)
accessTokenRes := c.GetAuthToken()

	/*----------------------------------------------个人用户管理模块-----------------------------------------------------*/

	//添加个人用户
	userAddReq := &requestModel.UserAddReq{ClientUserId: "1448519567", ClientUserName: "张志伟"}
	userAddResult := c.GetUserAddResponse(userAddReq, accessTokenRes.Data.AccessToken)
	fmt.Println(userAddResult)

	//更新个人用户
	userUpdateReq := &requestModel.UserUpdateReq{OpenUserId: "e1cb27058c68426cab534943f64dfbca",ClientUserId: "1448519567",ClientUserName: "张志伟22222"}
	userUpdateResult := c.GetUserUpdateResponse(userUpdateReq,accessTokenRes.Data.AccessToken)
	fmt.Println(userUpdateResult)

	//删除个人用户
	userDeleteReq := &requestModel.UserDeleteReq{OpenUserId: ""}
	userDeleteResult := c.GetUserDeleteResponse(userDeleteReq,accessTokenRes.Data.AccessToken)
	fmt.Println(userDeleteResult)

	//禁用个人用户
	userDisableReq := &requestModel.UserDisableReq{OpenUserId: "e1cb27058c68426cab534943f64dfbca"}
	userDisableResult := c.GetUserDisableResponse(userDisableReq,accessTokenRes.Data.AccessToken)
	fmt.Println(userDisableResult)

	//恢复个人用户
	userEnableReq := &requestModel.UserEnableReq{OpenUserId: "e1cb27058c68426cab534943f64dfbca"}
	userEnableResult := c.GetUserEnableResponse(userEnableReq,accessTokenRes.Data.AccessToken)
	fmt.Println(userEnableResult)

	//查询个人基本信息
	userInfoReq := &requestModel.GetUserReq{OpenUserId: "e1cb27058c68426cab534943f64dfbca"}
	userInfoResult := c.GetUserInfoResponse(userInfoReq,accessTokenRes.Data.AccessToken)
	fmt.Println(userInfoResult)

	//获取个人用户身份信息
	userIdentityReq := &requestModel.GetUserIdentifyInfoReq{OpenUserId: "e1cb27058c68426cab534943f64dfbca"}
	userIdentityResult := c.GetUserIdentityResponse(userIdentityReq,accessTokenRes.Data.AccessToken)
	fmt.Println(userIdentityResult)

	//获取个人实名认证进度
	userIdentityProcessReq := &requestModel.GetUserIdentityProgressReq{OpenUserId: "e1cb27058c68426cab534943f64dfbca"}
	userIdentityProcessResult := c.GetUserIdentityProcessResponse(userIdentityProcessReq,accessTokenRes.Data.AccessToken)
	fmt.Println(userIdentityProcessResult)

	//获取个人用户授权链接
	getUserAuthUrlReq := &requestModel.GetUserAuthUrlReq{OpenUserId: "e1cb27058c68426cab534943f64dfbca"}
	getUserAuthUrlRes := c.GetUserAuthUrlResponse(getUserAuthUrlReq, accessTokenRes.Data.AccessToken)
	fmt.Println(getUserAuthUrlRes)
	/*--------------------------------------------------企业用户管理模块-------------------------------------------------*/

	//添加企业用户
	corpAddReq := &requestModel.CorpAddReq{ClientCorpId: "1628479200230160177", ClientCorpName: "永诚保险"}
	corpAddRes := c.GetCorpAddResponse(corpAddReq, accessTokenRes.Data.AccessToken)
	fmt.Println(corpAddRes)

	//更新企业用户
	corpUpdateReq := &requestModel.CorpUpdateReq{OpenCorpId: "9b4f259d66ec435daae51b2d6f8d8179",ClientCorpId: "2222",ClientCorpName: "永诚保险"}
	corpUpdateResult := c.GetCorpUpdateResponse(corpUpdateReq,accessTokenRes.Data.AccessToken)
	fmt.Println(corpUpdateResult)

	//删除企业用户
	corpDeleteReq := &requestModel.CorpDeleteReq{OpenCorpId: ""}
	corpDeleteResult := c.GetCorpDeleteResponse(corpDeleteReq,accessTokenRes.Data.AccessToken)
	fmt.Println(corpDeleteResult)

	//禁用企业用户
	corpDisableReq := &requestModel.CorpDisableReq{OpenCorpId: "9b4f259d66ec435daae51b2d6f8d8179"}
	corpDisableRes := c.GetCorpDisableResponse(corpDisableReq,accessTokenRes.Data.AccessToken)
	fmt.Println(corpDisableRes)

	//恢复企业用户
	corpEnableReq := &requestModel.CorpEnableReq{OpenCorpId: "9b4f259d66ec435daae51b2d6f8d8179"}
	corpEnableRes := c.GetCorpEnableResponse(corpEnableReq,accessTokenRes.Data.AccessToken)
	fmt.Println(corpEnableRes)

	//查询企业用户信息
	corpInfoReq := &requestModel.GetCorpReq{OpenCorpId: "9b4f259d66ec435daae51b2d6f8d8179"}
	corpInfoRes := c.GetCorpInfoResponse(corpInfoReq,accessTokenRes.Data.AccessToken)
	fmt.Println(corpInfoRes)

	//查询企业用户身份信息
	corpIdentityReq := &requestModel.GetCorpIdentityInfoReq{OpenCorpId: "9b4f259d66ec435daae51b2d6f8d8179"}
	corpIdentityRes := c.GetCorpIdentityResponse(corpIdentityReq,accessTokenRes.Data.AccessToken)
	fmt.Println(corpIdentityRes)

	//查询企业用户实名认证进度
	corpIdentityProcessReq := &requestModel.GetCorpIdentityProgressReq{OpenCorpId: "9b4f259d66ec435daae51b2d6f8d8179"}
	corpIdentityProcessRes := c.GetCorpIdentityProcessResponse(corpIdentityProcessReq,accessTokenRes.Data.AccessToken)
	fmt.Println(corpIdentityProcessRes)

	//获取企业用户授权链接
	getCorpAuthUrlReq := &requestModel.GetCorpAuthUrlReq{OpenCorpId: "9b4f259d66ec435daae51b2d6f8d8179", OperatorId: "367999b2a50f411e8255396d583d3fce"}
	getCorpAuthUrlRes := c.GetCorpAuthUrlResponse(getCorpAuthUrlReq, accessTokenRes.Data.AccessToken)
	fmt.Println(getCorpAuthUrlRes)

	/*--------------------------------------------------企业链接管理模块-------------------------------------------------*/
	corpAuthUrlReq := &requestModel.GetCorpManageUrlReq{OpenCorpId: "9b4f259d66ec435daae51b2d6f8d8179",OperatorId: "a1d21b1ddbfa496dbf6cdab2c854e6a0",ResourceId: "seal"}
	corpAuthUrlRes := c.GetCorpManageUrlResponse(corpAuthUrlReq,accessTokenRes.Data.AccessToken)
	fmt.Println(corpAuthUrlRes)

	/*-------------------------------------------------文档和模板管理模块-------------------------------------------------*/

	//上传即时文件
	path := "C:\\Users\\Fadada\\Desktop\\FASC-OpenAPI V5.0.4.pdf"
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("本地文件读取失败")
	}
	ownerId := &commonModel.OpenId{IdType: "person", OpenUserId: "a1d21b1ddbfa496dbf6cdab2c854e6a0"}
	uploadFileReq := &requestModel.FileUploadInstantReq{OwnerId: ownerId, FileName: "FASC-OpenAPI V5.0.4.pdf", FileType: "attach", FileContentHash: common.GetFileHash(path)}
	uploadFileRes := c.UploadFileResponse(file, uploadFileReq, accessTokenRes.Data.AccessToken)
	fmt.Println(uploadFileRes.Data.FileId)
	
	//查询文档模板列表
	ownerId = &commonModel.OpenId{
		IdType:     "person",
		//OpenCorpId: "",
		OpenUserId: "a1d21b1ddbfa496dbf6cdab2c854e6a0",
	}
	//docListFilter := &requestModel.DocTemplateLisFilter{DocTemplateName: ""}
	docTemplateListReq := &requestModel.GetDocTemplateListReq{
		//OwnerId:      ownerId,
		//ListFilter:   docListFilter,
		//ListPageNo:   1,
		//ListPageSize: 1,
	}
	docTemplateListRes := c.GetDocTemplateListResponse(docTemplateListReq,accessTokenRes.Data.AccessToken)
	fmt.Println(docTemplateListRes)
	
	//查询模板详情
	docTemplateDetailReq := &requestModel.GetDocTemplateDetailReq{
		//OwnerId:       ownerId,
		DocTemplateId: "1635755968000131405",
	}
	docTemplateDetailRes := c.GetDocTemplateDetailResponse(docTemplateDetailReq,accessTokenRes.Data.AccessToken)
	fmt.Println(docTemplateDetailRes)

	//获取签署模拟列表
	//1.(如果需要参数)示例

// openId := &commonModel.OpenId{IdType: "person",OpenUserId:"8a3cb72ef85f4f62a6e82935eb41908c" } //signListFilter :=
&requestModel.SignTemplateListFilterInfo{SignTemplateName: "hangman"} getSignTemplateListReq :=
&requestModel.GetSignTemplateListReq{ //OwnerId:openId, //ListFilter: signListFilter, //ListPageNo: 1, // ListPageSize:
1, } getSignTemplateListRes := c.GetSignTemplateListResponse(getSignTemplateListReq, accessTokenRes.Data.AccessToken)
fmt.Println(getSignTemplateListRes)

	//获取签署模板详情
	signTemplateDetailReq := &requestModel.GetSignTemplateDetailReq{
		//OwnerId:        ownerId,
		SignTemplateId: "1629172127384170717",
	}
	signTemplateDetailRes := c.GetSignTemplateDetailResponse(signTemplateDetailReq,accessTokenRes.Data.AccessToken)
    fmt.Println(signTemplateDetailRes)

	//获取模板管理链接
	pageManageUrlReq := &requestModel.GetPageManageUrlReq{
		OpenCorpId:  "9b4f259d66ec435daae51b2d6f8d8179",
		//RedirectUrl: "",
	}
	pageManageUrlRes := c.GetPageManageUrlResponse(pageManageUrlReq,accessTokenRes.Data.AccessToken)
	fmt.Printf("获取模板链接接口请求结果%s",pageManageUrlRes)
	/*-------------------------------------------------签署任务管理模块--------------------------------------------------*/

	//创建签署任务
	signTaskDocInfo1 := &commonModel.SignTaskDocInfo{
		DocId:         2,
		DocName:       "保全报告1314",
		//DocFileId:     "",
		DocTemplateId: "1635755968000131405",
		//DocFields:     nil,
	}
	docInfoArr := &[]commonModel.SignTaskDocInfo{*signTaskDocInfo1}
	//创建签署任务
	signTaskReq := &requestModel.CreateSignTaskReq{
		SignTaskSubject: "基于go版本的sdk开始测试",
		Initiator:       ownerId,
		//Docs:            docInfoArr,
	}
	signTaskRes := c.CreateSignTask(signTaskReq,accessTokenRes.Data.AccessToken)
	fmt.Println(signTaskRes)


	//创建签署任务基于签署模板
	initiator := &commonModel.OpenId{IdType: "person", OpenUserId: "8a3cb72ef85f4f62a6e82935eb41908c"}
	// businessScene := &requestModel.BusinessSceneInfo{BusinessId: ""} //(需要时再添加)

	notification := &commonModel.Notification{SendNotification: false}
	fillActor := &commonModel.Actor{ActorType: "filler", ActorId: "111", ActorIdentType: "person", Notification: notification}

	actorField1 := &commonModel.FieldValueInfo{FieldDocId: 111, FieldValue: ""}
	actorField2 := &commonModel.FieldValueInfo{FieldDocId: 222, FieldValue: ""}
	actorFieldsArr := &[]commonModel.FieldValueInfo{*actorField1, *actorField2}

	fillActor1 := &commonModel.FillActorInfo{FillActor: fillActor, ActorFields: actorFieldsArr}
	fillActorArr := &[]commonModel.FillActorInfo{*fillActor1}

	createSignTaskWithTemReq := &requestModel.CreateSignTaskWithTemplateReq{SignTaskSubject: "基于golang创建签署任务", SignTemplateId: "1635243428580177021", Initiator: initiator, FillActors: fillActorArr}
	createSignTaskWithTemRes := c.CreateSignTaskWithTemplate(createSignTaskWithTemReq, accessTokenRes.Data.AccessToken)
	fmt.Println(createSignTaskWithTemRes)

	//添加签署任务文档
	addDocInfoReq := &requestModel.AddSignTaskDocReq{
		SignTaskId: "1646202825361115359",
		Docs:       docInfoArr,
	}
	addDocInfoRes := c.AddSignTaskDoc(addDocInfoReq,accessTokenRes.Data.AccessToken)
	fmt.Println(addDocInfoRes)

	//移除签署任务文档
	docIds := []int{1}
	deleteDocInfoReq := &requestModel.DeleteSignTaskDocReq{
		SignTaskId: "1646202825361115359",
		DocIds: docIds,
	}
	deleteDocInfoRes := c.DeleteSignTaskDoc(deleteDocInfoReq,accessTokenRes.Data.AccessToken)
	fmt.Println(deleteDocInfoRes)

	//添加签署任务控件
	field1 := &commonModel.Field{
		FieldId:             "1314520",
		FieldName:           "企业印章",
		Position:            &commonModel.FieldPosition{
			PositionMode:    "pixel",
			PositionPageNo:  1,
			PositionX:       "200",
			PositionY:       "300",
		},
		FieldType:           "corp_seal",
	//	FieldTextSingleLine: nil,
		//FieldTextMultiLine:  nil,
		//FieldCheckBox:       nil,
	}
	field2 := &commonModel.Field{
		FieldId:             "5201314",
		FieldName:           "企业骑缝章",
		Position:             &commonModel.FieldPosition{
			PositionMode:    "pixel",
			PositionPageNo:  1,
			PositionX:       "400",
			PositionY:       "300",
		},
		FieldType:           "corp_seal_cross_page",
		//FieldTextSingleLine: nil,
		//FieldTextMultiLine:  nil,
		//FieldCheckBox:       nil,
	}
	fieldArr1 := &[]commonModel.Field{*field2,*field1}
	signTaskField1 := &requestModel.AddSignTaskField{
		DocId:     2,
		DocFields: fieldArr1,
	}
	docFieldArr := &[]requestModel.AddSignTaskField{*signTaskField1}
	addFieldReq := &requestModel.AddSignTaskFieldReq{
		SignTaskId: "1646202825361115359",
		Fields:     docFieldArr,
	}
	addFieldRes := c.AddSignTaskField(addFieldReq,accessTokenRes.Data.AccessToken)
	fmt.Println(addFieldRes)

	//移除签署任务控件
	docFields1 := []string{"1","2"}
	delSignTaskField1 := &requestModel.DeleteSignTaskField{
		DocId:     2,
		DocFields: docFields1,
	}
    deleteDocFieldArr := &[]requestModel.DeleteSignTaskField{*delSignTaskField1}
	deleteFieldReq := &requestModel.DeleteSignTaskFieldReq{
		SignTaskId: "",
		Fields:     deleteDocFieldArr,
	}
	deleteFieldRes := c.DeleteSignTaskField(deleteFieldReq,accessTokenRes.Data.AccessToken)
    fmt.Println(deleteFieldRes)

	//获取签署任务控制设置链接
	//userIdentityInfo := &commonModel.UserIdentInfo{
		//UserName:  "",
		//IdentType: "",
		//IdentNo:   "",
	//}
	//userInfoExtend := &commonModel.UserInfoExtend{
	//	BankAccountNo: "",
		//Mobile:        "",

// } signTaskFieldUrlReq := &requestModel.GetSignTaskFieldUrlReq{ SignTaskId:     "1646202825361115359",
OpenUserId:     "a1d21b1ddbfa496dbf6cdab2c854e6a0", //UserIdentInfo:  userIdentityInfo, //UserInfoExtend:
userInfoExtend, //RedirectUrl:    "", } signTaskFieldUrlRes := c.GetSignTaskFieldUrl(
signTaskFieldUrlReq,accessTokenRes.Data.AccessToken)
fmt.Println(signTaskFieldUrlRes)

	//获取签署任务详情
	signTaskDetailReq := &requestModel.GetSignTaskDetailReq{SignTaskId: "1646202825361115359"}
	signTaskDetailRes := c.GetSignTaskDetailResponse(signTaskDetailReq, accessTokenRes.Data.AccessToken)
	fmt.Println(signTaskDetailRes)

	//填写签署任务控件内容
	fieldValues1 := &requestModel.DocFieldValue{
		DocId:      0,
		FieldId:    "0001",
		FieldValue: "填写控件内容1",
	}
	fillFieldValuesReq := &requestModel.FillFieldValuesReq{
		SignTaskId:     "1644903847541180130",
		DocFieldValues: &[]requestModel.DocFieldValue{*fieldValues1},
	}
	fillFieldValuesRes := c.FillSignTaskFieldValues(fillFieldValuesReq, accessTokenRes.Data.AccessToken)
	fmt.Println(fillFieldValuesRes)

	//添加签署任务附件
	attach1 := &commonModel.SignTaskAttachs{
		AttachId:     10,
		AttachName:   "咚咚咚",
		AttachFileId: "1646206567814148150",
	}
	addAttachReq := &requestModel.AddSignTaskAttachReq{
		SignTaskId: "",
		Attachs:   &[]commonModel.SignTaskAttachs{*attach1},
	}
	addAttachRes := c.AddSignTaskAttach(addAttachReq,accessTokenRes.Data.AccessToken)
	fmt.Println(addAttachRes)
	
	//移除签署任务附件
	deleteAttachReq := &requestModel.DeleteSignTaskAttachReq{
		SignTaskId: "1646202825361115359",
		AttachIds:  []int{10},
	}
	deleteAttachRes := c.DeleteSignTaskAttach(deleteAttachReq,accessTokenRes.Data.AccessToken)
	fmt.Println(deleteAttachRes)

	//添加签署任务参与方
     //填写方列表，如果需要添加签署方以及抄送方列表写法类似
	signfillActor1 := &requestModel.SignTaskFillActor{
		FillActor : &commonModel.Actor{
			ActorType:      "filler",
			ActorId:        "个人",
			ActorIdentType: "person",
			//ActorUser:      nil,
			//ActorCorp:      nil,
			Notification:   &commonModel.Notification{
				SendNotification: false,
				//NotifyWay:        "",
				//NotifyAddress:    "",
			},
		},
	}
	signFillActors := &[]requestModel.SignTaskFillActor{*signfillActor1}
	addActorsReq := &requestModel.AddSignTaskActorReq{
		SignTaskId: "1645581657578114293",
		FillActors: signFillActors,
		//SignActors: nil,
		//CcActors:   nil,
	}
	addActorsRes := c.AddSignTaskActor(addActorsReq,accessTokenRes.Data.AccessToken)
	fmt.Println(addActorsRes)

	//移除签署任务参与方
	deleteActorsReq := &requestModel.DeleteSignTaskActorReq{
		SignTaskId: "",
		ActorType:  "signer",
		ActorIds:   []string{"个人"},
	}
	deleteActorsRes := c.DeleteSignTaskActor(deleteActorsReq,accessTokenRes.Data.AccessToken)
	fmt.Println(deleteActorsRes)

	//发起签署任务
	initiateSignTaskReq := &requestModel.InitiateSignTaskReq{SignTaskId: "1646202825361115359"}
	initiateSignTaskRes := c.InitiateSignTask(initiateSignTaskReq,accessTokenRes.Data.AccessToken)
	fmt.Println(initiateSignTaskRes)

	//撤销签署任务
	cancelSignTaskReq := &requestModel.CancelSignTaskReq{SignTaskId: ""}
	cancelSignTaskRes := c.CancelSignTask(cancelSignTaskReq,accessTokenRes.Data.AccessToken)
	fmt.Println(cancelSignTaskRes)

	//定稿签署任务文档
	finalizeSignTaskReq := &requestModel.FinalizeSignTaskReq{SignTaskId: "1644903847541180130"}
	finalizeSignTaskRes := c.FinalizeSignTaskDoc(finalizeSignTaskReq,accessTokenRes.Data.AccessToken)
	fmt.Println(finalizeSignTaskRes)

	//阻塞签署任务
	blockSignTaskReq := &requestModel.BlockSignTaskReq{
		SignTaskId: "1644903847541180130",
		ActorType:  "signer",
		ActorId:    "个人",
	}
	blockSignTaskRes := c.BlockSignTask(blockSignTaskReq,accessTokenRes.Data.AccessToken)
	fmt.Println(blockSignTaskRes)

	//解阻签署任务
	unBlockSignTaskReq := &requestModel.UnBlockSignTaskReq{
		SignTaskId: "1644903847541180130",
		ActorType:  "signer",
		ActorId:    "个人",
	}
	unBlockSignTaskRes := c.UnBlockSignTask(unBlockSignTaskReq,accessTokenRes.Data.AccessToken)
	fmt.Println(unBlockSignTaskRes)

	//催办签署任务
	urgeSignTaskReq := &requestModel.UrgeSignTaskReq{SignTaskId: "1644551524323139545"}
	urgeSignTaskRes := c.UrgeSignTask(urgeSignTaskReq,accessTokenRes.Data.AccessToken)
	fmt.Printf("催办签署任务请求结果%s",urgeSignTaskRes)

	//结束签署任务
	finishSignTaskReq := &requestModel.FinishSignTaskReq{SignTaskId: "1644903847541180130"}
	finishSignTaskRes := c.FinishSignTask(finishSignTaskReq,accessTokenRes.Data.AccessToken)
	fmt.Println(finishSignTaskRes)

	//获取签署任务链接
	signTaskUrlReq := &requestModel.GetSignTaskUrlReq{SignTaskId: "1644551524323139545", ActorType: "initiator"}
	signTaskUrlRes := c.GetSignTaskUrlResponse(signTaskUrlReq, accessTokenRes.Data.AccessToken)
	fmt.Println(signTaskUrlRes)

	//下载签署任务文档
	downloadFilesReq := &requestModel.DownloadFilesReq{SignTaskId: "1640935237063137559"}
	downloadFilesRes := c.DownloadSignTaskFiles(downloadFilesReq, accessTokenRes.Data.AccessToken)

	if find := strings.Contains(downloadFilesRes.ContentType, ContentTypeJson); find {
		fmt.Println(downloadFilesRes.Data)
	}
	if find := strings.Contains(downloadFilesRes.ContentType, ContentTypeZip); find {
		file, err := os.Create("D:\\downLoads\\" +downloadFilesReq.SignTaskId +".zip")//用signTaskId命名最好
		if err != nil {
		}
		_, err = file.Write(downloadFilesRes.Content)
		if err != nil {
			fmt.Println("文件写入失败")
		}
	}
	if find := strings.Contains(downloadFilesRes.ContentType, ContentTypePdf); find {
		file, err := os.Create("D:\\downLoads\\" +downloadFilesReq.SignTaskId +".pdf")
		if err != nil {
		}
		_, err = file.Write(downloadFilesRes.Content)
		if err != nil {
			fmt.Println("文件写入失败")
		}
	}
	if find := strings.Contains(downloadFilesRes.ContentType, ContentTypeFile); find {
		file, err := os.Create("D:\\downLoads\\" +downloadFilesReq.SignTaskId +"")//根据已知的文件类型去定义文件后缀
		if err != nil {
		}
		_, err = file.Write(downloadFilesRes.Content)
		if err != nil {
			fmt.Println("文件写入失败")
		}
	}

	/*---------------------------------------------------计费管理模块---------------------------------------------------*/

	//获取计费链接
	billUrlReq := &requestModel.GetBillingUrlReq{
		OpenId:      ownerId,
		//UrlType:     "",
		//RedirectUrl: "",
	}
	billUrlRes := c.GetBillUrlResponse(billUrlReq,accessTokenRes.Data.AccessToken)
	fmt.Println(billUrlRes)

-----更多调用示例敬请期待


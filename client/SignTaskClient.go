package client

import (
	"encoding/json"
	"fmt"
	common2 "github.com/ZZWisking/fasc_openapi_go_sdk/common"
	requestModel "github.com/ZZWisking/fasc_openapi_go_sdk/model/requestModel"
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/responseModel"
	"strings"
)

const (
	CreateSignTaskReqUrlWithTem   = "/sign-task/create-with-template" //创建签署任务基于签署模板
	GetSignTaskDetailReqUrl       = "/sign-task/get-detail"           //获取签署任务详情
	GetSignTaskUrlReqUrl          = "/sign-task/get-url"              //获取签署任务链接
	DownSignTaskFileReqUrl        = "/sign-task/download-files"       //下载签署任务文档
	CreateSignTaskReqUrl          = "/sign-task/create"               //创建签署任务
	AddSignTaskDocReqUrl          = "/sign-task/doc/add"              //添加签署任务文档
	DeleteSignTaskDocReqUrl       = "/sign-task/doc/delete"           //移除签署任务文档
	AddSignTaskFieldReqUrl        = "/sign-task/field/add"            //添加签署任务控件
	DeleteSignTaskFieldReqUrl     = "/sign-task/field/delete"         //移除签署任务控件
	GetSignTaskFieldReqUrl        = "/sign-task/field/get-url"        //获取签署任务控件设置链接
	FillSignTaskFieldValuesReqUrl = "/sign-task/field/fill-values"    //填写签署任务控件内容
	AddSignTaskAttachReqUrl       = "/sign-task/attach/add"           //添加签署任务附件
	DeleteSignTaskAttachReqUrl    = "/sign-task/attach/delete"        //移除签署任务附件
	AddSignTaskActorReqUrl        = "/sign-task/actor/add"            //添加签署任务参与方
	DeleteSignTaskActorReqUrl     = "/sign-task/actor/delete"         //移除签署任务参与方
	InitiateSignTaskReqUrl        = "/sign-task/initiate"             //发起签署任务
	CancelSignTaskReqUrl          = "/sign-task/cancel"               //撤销签署任务
	FinalizeSignTaskDocReqUrl     = "/sign-task/doc/finalize"         //定稿签署任务文档
	BlockSignTaskReqUrl           = "/sign-task/block"                //阻塞签署任务
	UnblockSignTaskReqUrl         = "/sign-task/unblock"              //解阻签署任务
	FinishSignTaskReqUrl          = "/sign-task/finish"               //结束签署任务
	ContentType                   = "application/json"
)

// CreateSignTask 创建签署任务
func (o *openApiClient) CreateSignTask(createSignTaskReq *requestModel.CreateSignTaskReq, accessToken string) responseModel.CreateSignTaskRes {
	var createSignTaskRes responseModel.CreateSignTaskRes
	reqStr, err := json.Marshal(createSignTaskReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + CreateSignTaskReqUrl                             //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &createSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	createSignTaskRes.RequestId = requestId
	return createSignTaskRes
}

// CreateSignTaskWithTemplate 创建签署任务基于签署模板
func (o *openApiClient) CreateSignTaskWithTemplate(addSignTaskReq *requestModel.CreateSignTaskWithTemplateReq, accessToken string) responseModel.CreateSignTaskWithTemplateRes {
	var addSignTaskRes responseModel.CreateSignTaskWithTemplateRes
	reqStr, err := json.Marshal(addSignTaskReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + CreateSignTaskReqUrlWithTem                      //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &addSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	addSignTaskRes.RequestId = requestId
	return addSignTaskRes
}

// AddSignTaskDoc 添加签署任务文档
func (o *openApiClient) AddSignTaskDoc(addSignTaskDocReq *requestModel.AddSignTaskDocReq, accessToken string) responseModel.AddSignTaskDocRes {
	var addSignTaskDocRes responseModel.AddSignTaskDocRes
	reqStr, err := json.Marshal(addSignTaskDocReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + AddSignTaskDocReqUrl                             //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &addSignTaskDocRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	addSignTaskDocRes.RequestId = requestId
	return addSignTaskDocRes
}

// DeleteSignTaskDoc 移除签署任务文档
func (o *openApiClient) DeleteSignTaskDoc(deleteSignTaskDocReq *requestModel.DeleteSignTaskDocReq, accessToken string) responseModel.DeleteSignTaskDocRes {
	var deleteSignTaskDocRes responseModel.DeleteSignTaskDocRes
	reqStr, err := json.Marshal(deleteSignTaskDocReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + DeleteSignTaskDocReqUrl                          //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &deleteSignTaskDocRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	deleteSignTaskDocRes.RequestId = requestId
	return deleteSignTaskDocRes
}

// AddSignTaskField 添加签署任务控件
func (o *openApiClient) AddSignTaskField(addSignTaskFieldReq *requestModel.AddSignTaskFieldReq, accessToken string) responseModel.AddSignTaskFieldRes {
	var addSignTaskFieldRes responseModel.AddSignTaskFieldRes
	reqStr, err := json.Marshal(addSignTaskFieldReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + AddSignTaskFieldReqUrl                           //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &addSignTaskFieldRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	addSignTaskFieldRes.RequestId = requestId
	return addSignTaskFieldRes
}

// DeleteSignTaskField 移除签署任务控件
func (o *openApiClient) DeleteSignTaskField(deleteSignTaskFieldReq *requestModel.DeleteSignTaskFieldReq, accessToken string) responseModel.DeleteSignTaskFieldRes {
	var deleteSignTaskFieldRes responseModel.DeleteSignTaskFieldRes
	reqStr, err := json.Marshal(deleteSignTaskFieldReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + DeleteSignTaskFieldReqUrl                        //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &deleteSignTaskFieldRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	deleteSignTaskFieldRes.RequestId = requestId
	return deleteSignTaskFieldRes
}

// GetSignTaskFieldUrl 获取签署任务控件设置链接
func (o *openApiClient) GetSignTaskFieldUrl(signTaskFieldUrlReq *requestModel.GetSignTaskFieldUrlReq, accessToken string) responseModel.GetSignTaskFieldUrlRes {
	var signTaskFieldUrlRes responseModel.GetSignTaskFieldUrlRes
	reqStr, err := json.Marshal(signTaskFieldUrlReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetSignTaskFieldReqUrl                           //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &signTaskFieldUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	signTaskFieldUrlRes.RequestId = requestId
	return signTaskFieldUrlRes
}

// FillSignTaskFieldValues 填写签署任务控件内容
func (o *openApiClient) FillSignTaskFieldValues(fillFieldReq *requestModel.FillFieldValuesReq, accessToken string) responseModel.FillFieldValuesRes {
	var fillFieldRes responseModel.FillFieldValuesRes
	reqStr, err := json.Marshal(fillFieldReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + FillSignTaskFieldValuesReqUrl                    //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &fillFieldRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	fillFieldRes.RequestId = requestId
	return fillFieldRes
}

// AddSignTaskAttach 添加签署任务附件
func (o *openApiClient) AddSignTaskAttach(addAttachReq *requestModel.AddSignTaskAttachReq, accessToken string) responseModel.AddSignTaskAttachRes {
	var addAttachRes responseModel.AddSignTaskAttachRes
	reqStr, err := json.Marshal(addAttachReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + AddSignTaskAttachReqUrl                          //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &addAttachRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	addAttachRes.RequestId = requestId
	return addAttachRes
}

// DeleteSignTaskAttach 移除签署任务附件
func (o *openApiClient) DeleteSignTaskAttach(deleteAttachReq *requestModel.DeleteSignTaskAttachReq, accessToken string) responseModel.DeleteSignTaskAttachRes {
	var deleteAttachRes responseModel.DeleteSignTaskAttachRes
	reqStr, err := json.Marshal(deleteAttachReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + DeleteSignTaskAttachReqUrl                       //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &deleteAttachRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	deleteAttachRes.RequestId = requestId
	return deleteAttachRes
}

// AddSignTaskActor 添加签署任务参与方
func (o *openApiClient) AddSignTaskActor(addActorReq *requestModel.AddSignTaskActorReq, accessToken string) responseModel.AddSignTaskActorRes {
	var addActorRes responseModel.AddSignTaskActorRes
	reqStr, err := json.Marshal(addActorReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + AddSignTaskActorReqUrl                           //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &addActorRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	addActorRes.RequestId = requestId
	return addActorRes
}

// DeleteSignTaskActor 移除签署任务参与方
func (o *openApiClient) DeleteSignTaskActor(deleteActorReq *requestModel.DeleteSignTaskActorReq, accessToken string) responseModel.DeleteSignTaskActorRes {
	var deleteActorRes responseModel.DeleteSignTaskActorRes
	reqStr, err := json.Marshal(deleteActorReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + DeleteSignTaskActorReqUrl                        //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &deleteActorRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	deleteActorRes.RequestId = requestId
	return deleteActorRes
}

// InitiateSignTask 发起签署任务
func (o *openApiClient) InitiateSignTask(initiateSignTaskReq *requestModel.InitiateSignTaskReq, accessToken string) responseModel.InitiateSignTaskRes {
	var initiateSignTaskRes responseModel.InitiateSignTaskRes
	reqStr, err := json.Marshal(initiateSignTaskReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + InitiateSignTaskReqUrl                           //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &initiateSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	initiateSignTaskRes.RequestId = requestId
	return initiateSignTaskRes
}

// CancelSignTask 撤销签署任务
func (o *openApiClient) CancelSignTask(cancelSignTaskReq *requestModel.CancelSignTaskReq, accessToken string) responseModel.CancelSignTaskRes {
	var cancelSignTaskRes responseModel.CancelSignTaskRes
	reqStr, err := json.Marshal(cancelSignTaskReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + CancelSignTaskReqUrl                             //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &cancelSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	cancelSignTaskRes.RequestId = requestId
	return cancelSignTaskRes
}

// FinalizeSignTaskDoc 定稿签署任务文档
func (o *openApiClient) FinalizeSignTaskDoc(finalizeDocReq *requestModel.FinalizeSignTaskReq, accessToken string) responseModel.FinalizeSignTaskRes {
	var finalizeDocRes responseModel.FinalizeSignTaskRes
	reqStr, err := json.Marshal(finalizeDocReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + FinishSignTaskReqUrl                             //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &finalizeDocRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	finalizeDocRes.RequestId = requestId
	return finalizeDocRes
}

// BlockSignTask 阻塞签署任务
func (o *openApiClient) BlockSignTask(blockSignTaskReq *requestModel.BlockSignTaskReq, accessToken string) responseModel.BlockSignTaskRes {
	var blockSignTaskRes responseModel.BlockSignTaskRes
	reqStr, err := json.Marshal(blockSignTaskReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + BlockSignTaskReqUrl                              //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &blockSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	blockSignTaskRes.RequestId = requestId
	return blockSignTaskRes
}

// UnBlockSignTask 解阻签署任务
func (o *openApiClient) UnBlockSignTask(unBlockSignTaskReq *requestModel.UnBlockSignTaskReq, accessToken string) responseModel.UnBlockSignTaskRes {
	var unBlockSignTaskRes responseModel.UnBlockSignTaskRes
	reqStr, err := json.Marshal(unBlockSignTaskReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + UnblockSignTaskReqUrl                            //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &unBlockSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	unBlockSignTaskRes.RequestId = requestId
	return unBlockSignTaskRes
}

// FinishSignTask 结束签署任务
func (o *openApiClient) FinishSignTask(finishSignTaskReq *requestModel.FinishSignTaskReq, accessToken string) responseModel.FinishSignTaskRes {
	var finishSignTaskRes responseModel.FinishSignTaskRes
	reqStr, err := json.Marshal(finishSignTaskReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + FinishSignTaskReqUrl                             //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &finishSignTaskRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	finishSignTaskRes.RequestId = requestId
	return finishSignTaskRes
}

// GetSignTaskDetailResponse 获取签署任务详情
func (o *openApiClient) GetSignTaskDetailResponse(signTaskDetailReq *requestModel.GetSignTaskDetailReq, accessToken string) responseModel.GetSignTaskDetailRes {
	var signTaskDetailRes responseModel.GetSignTaskDetailRes
	reqStr, err := json.Marshal(signTaskDetailReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetSignTaskDetailReqUrl                          //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &signTaskDetailRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	signTaskDetailRes.RequestId = requestId
	return signTaskDetailRes
}

// GetSignTaskUrlResponse 获取签署任务链接
func (o *openApiClient) GetSignTaskUrlResponse(signTaskUrlReq *requestModel.GetSignTaskUrlReq, accessToken string) responseModel.GetSignTaskUrlRes {
	var signTaskUrlRes responseModel.GetSignTaskUrlRes
	reqStr, err := json.Marshal(signTaskUrlReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetSignTaskUrlReqUrl                             //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &signTaskUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	signTaskUrlRes.RequestId = requestId
	return signTaskUrlRes
}

// DownloadSignTaskFiles 下载签署任务文档
func (o *openApiClient) DownloadSignTaskFiles(downloadFilesReq *requestModel.DownloadFilesReq, accessToken string) responseModel.DownloadFilesRes {
	var downLoadFilesRes responseModel.DownloadFilesRes
	reqStr, err := json.Marshal(downloadFilesReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                                                     //拼接post提交body参数
	requestUrl := o.serverUrl + DownSignTaskFileReqUrl                                                 //接口请求api地址
	rspBody, requestId, contentType := common2.DownLoadFilesPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	downLoadFilesRes.RequestId = requestId
	downLoadFilesRes.Content = rspBody
	downLoadFilesRes.ContentType = contentType
	if find := strings.Contains(contentType, ContentType); find {
		downLoadFilesRes.Data = string(rspBody)
	}
	return downLoadFilesRes
}

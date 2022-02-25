package client

import (
	"encoding/json"
	common2 "github.com/ZZWisking/fasc_openapi_go_sdk/common"
	requestModel2 "github.com/ZZWisking/fasc_openapi_go_sdk/model/requestModel"
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/responseModel"
	"fmt"
	"strings"
)

const(
	CreateSignTaskReqUrlWithTem ="/sign-task/create-with-template" //创建签署任务基于签署模板
	GetSignTaskDetailReqUrl = "/sign-task/get-detail" //获取签署任务详情
	GetSignTaskUrlReqUrl = "/sign-task/get-url" //获取签署任务链接
	DownSignTaskFileReqUrl = "/sign-task/download-files" //下载签署任务文档
	ContentType = "application/json"
)

// CreateSignTaskWithTemplate 创建签署任务基于签署模板
func (o *openApiClient) CreateSignTaskWithTemplate(addSignTaskReq *requestModel2.CreateSignTaskWithTemplateReq,accessToken string) responseModel.CreateSignTaskWithTemplateRes  {
	var addSignTaskRes responseModel.CreateSignTaskWithTemplateRes
	reqStr , err := json.Marshal(addSignTaskReq)
	if err != nil{
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken,string(reqStr))
	bodyStr := "bizContent"+ "="+ string(reqStr)//拼接post提交body参数
	requestUrl := o.serverUrl + CreateSignTaskReqUrlWithTem         //接口请求api地址
	rspBody,requestId := common2.SendPost(requestUrl,headMap,[]byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody,&addSignTaskRes)
	if err != nil{
		fmt.Println("json字符串转为struct失败")
	}
	addSignTaskRes.RequestId = requestId
	return addSignTaskRes
}

// GetSignTaskDetailResponse 获取签署任务详情
func (o *openApiClient) GetSignTaskDetailResponse(signTaskDetailReq *requestModel2.GetSignTaskDetailReq,accessToken string) responseModel.GetSignTaskDetailRes  {
	var signTaskDetailRes responseModel.GetSignTaskDetailRes
	reqStr , err := json.Marshal(signTaskDetailReq)
	if err != nil{
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken,string(reqStr))
	bodyStr := "bizContent"+ "="+ string(reqStr)//拼接post提交body参数
	requestUrl := o.serverUrl + GetSignTaskDetailReqUrl         //接口请求api地址
	rspBody,requestId := common2.SendPost(requestUrl,headMap,[]byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody,&signTaskDetailRes)
	if err != nil{
		fmt.Println("json字符串转为struct失败")
	}
	signTaskDetailRes.RequestId = requestId
	return signTaskDetailRes
}

// GetSignTaskUrlResponse 获取签署任务链接
func (o *openApiClient) GetSignTaskUrlResponse(signTaskUrlReq *requestModel2.GetSignTaskUrlReq,accessToken string) responseModel.GetSignTaskUrlRes  {
	var signTaskUrlRes responseModel.GetSignTaskUrlRes
	reqStr , err := json.Marshal(signTaskUrlReq)
	if err != nil{
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken,string(reqStr))
	bodyStr := "bizContent"+ "="+ string(reqStr)//拼接post提交body参数
	requestUrl := o.serverUrl + GetSignTaskUrlReqUrl         //接口请求api地址
	rspBody,requestId := common2.SendPost(requestUrl,headMap,[]byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody,&signTaskUrlRes)
	if err != nil{
		fmt.Println("json字符串转为struct失败")
	}
	signTaskUrlRes.RequestId = requestId
	return signTaskUrlRes
}

// DownloadSignTaskFiles 下载签署任务文档
func (o *openApiClient) DownloadSignTaskFiles(downloadFilesReq *requestModel2.DownloadFilesReq,accessToken string) responseModel.DownloadFilesRes {
	var downLoadFilesRes responseModel.DownloadFilesRes
	reqStr , err := json.Marshal(downloadFilesReq)
	if err != nil{
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken,string(reqStr))
	bodyStr := "bizContent"+ "="+ string(reqStr)//拼接post提交body参数
	requestUrl := o.serverUrl + DownSignTaskFileReqUrl         //接口请求api地址
	rspBody,requestId,contentType := common2.DownLoadFilesPost(requestUrl,headMap,[]byte(bodyStr)) //发送post请求
	downLoadFilesRes.RequestId = requestId
	downLoadFilesRes.Content = rspBody
	downLoadFilesRes.ContentType = contentType
	if find := strings.Contains(contentType,ContentType); find{
		downLoadFilesRes.Data = string(rspBody)
	}
	return downLoadFilesRes
}

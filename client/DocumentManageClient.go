package client

import (
	"encoding/json"
	"fmt"
	common2 "github.com/ZZWisking/fasc_openapi_go_sdk/common"
	requestModel "github.com/ZZWisking/fasc_openapi_go_sdk/model/requestModel"
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/responseModel"
)

const (
	GetSignTemplateListReqUrl   = "/sign-template/get-list"   //获取签署模板列表
	FileUploadReqUrl            = "/file/upload-instant"      //上传即时文件
	GetDocTemplateListReqUrl    = "/doc-template/get-list"    //查询文档模板列表
	GetDocTemplateDetailReqUrl  = "/doc-template/get-detail"  //获取文档模板详情
	GetSignTemplateDetailReqUrl = "/sign-template/get-detail" //获取签署模板详情
)

// UploadFileResponse 上传即时文件
func (o *openApiClient) UploadFileResponse(file []byte, uploadFileReq *requestModel.FileUploadInstantReq, accessToken string) responseModel.FileUploadInstantRes {
	var uploadFileRes responseModel.FileUploadInstantRes
	reqStr, err := json.Marshal(uploadFileReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	requestUrl := o.serverUrl + FileUploadReqUrl                                                                         //接口请求api地址
	rspBody, requestId, err := common2.UploadFilePost(requestUrl, headMap, file, string(reqStr), uploadFileReq.FileName) //发送post请求
	if err != nil {
		fmt.Printf("文件上传请求发生异常,%s", err)
	}
	err = json.Unmarshal(rspBody, &uploadFileRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	uploadFileRes.RequestId = requestId
	return uploadFileRes
}

// GetSignTemplateListResponse 查询签署模板列表
func (o *openApiClient) GetSignTemplateListResponse(signTemplateReq *requestModel.GetSignTemplateListReq, accessToken string) responseModel.GetSignTemplateListRes {
	var signTemplateRes responseModel.GetSignTemplateListRes
	reqStr, err := json.Marshal(signTemplateReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetSignTemplateListReqUrl                        //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &signTemplateRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	signTemplateRes.RequestId = requestId
	return signTemplateRes
}

// GetSignTemplateDetailResponse 获取签署模板详情
func (o *openApiClient) GetSignTemplateDetailResponse(signTemDetailReq *requestModel.GetSignTemplateDetailReq, accessToken string) responseModel.GetSignTaskDetailRes {
	var signTemDetailRes responseModel.GetSignTaskDetailRes
	reqStr, err := json.Marshal(signTemDetailReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetSignTemplateDetailReqUrl                      //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &signTemDetailRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	signTemDetailRes.RequestId = requestId
	return signTemDetailRes
}

// GetDocTemplateDetailResponse 获取文档模板详情
func (o *openApiClient) GetDocTemplateDetailResponse(docTemDetailReq *requestModel.GetDocTemplateDetailReq, accessToken string) responseModel.GetDocTemplateDetailRes {
	var docTemDetailRes responseModel.GetDocTemplateDetailRes
	reqStr, err := json.Marshal(docTemDetailReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetDocTemplateDetailReqUrl                       //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &docTemDetailRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	docTemDetailRes.RequestId = requestId
	return docTemDetailRes
}

// GetDocTemplateListResponse 查询文档模板列表
func (o *openApiClient) GetDocTemplateListResponse(docTemListReq *requestModel.GetDocTemplateListReq, accessToken string) responseModel.GetDocTemplateListRes {
	var docTemListRes responseModel.GetDocTemplateListRes
	reqStr, err := json.Marshal(docTemListReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetDocTemplateListReqUrl                         //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &docTemListRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	docTemListRes.RequestId = requestId
	return docTemListRes
}

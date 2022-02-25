package client

import (
	"encoding/json"
	common2 "fasc_openapi_go_sdk/common"
	requestModel2 "fasc_openapi_go_sdk/model/requestModel"
	"fasc_openapi_go_sdk/model/responseModel"
	"fmt"
)

const(
	CorpAddReqUrl ="/corp/add" //添加企业用户
	GetCorpAuthReqUrl ="/corp/get-auth-url" //获取企业用户授权链接
)

// GetCorpAddResponse 添加企业用户
func (o *openApiClient) GetCorpAddResponse(corpAddReq *requestModel2.CorpAddReq,accessToken string) responseModel.CorpAddRes  {
	var corpAddRes responseModel.CorpAddRes
	reqStr , err := json.Marshal(corpAddReq)
	if err != nil{
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken,string(reqStr))
	bodyStr := "bizContent"+ "="+ string(reqStr)//拼接post提交body参数
	requestUrl := o.serverUrl + CorpAddReqUrl              //接口请求api地址
	rspBody,requestId := common2.SendPost(requestUrl,headMap,[]byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody,&corpAddRes)
	if err != nil{
		fmt.Println("json字符串转为struct失败")
	}
	corpAddRes.RequestId = requestId
	return corpAddRes
}

// GetCorpAuthUrlResponse 获取企业用户授权链接
func (o *openApiClient) GetCorpAuthUrlResponse(gerCorpAuthUrlReq *requestModel2.GetCorpAuthUrlReq,accessToken string) responseModel.GetCorpAuthUrlRes  {
	var gerCorpAuthUrlRes responseModel.GetCorpAuthUrlRes
	reqStr , err := json.Marshal(gerCorpAuthUrlReq)
	if err != nil{
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken,string(reqStr))
	bodyStr := "bizContent"+ "="+ string(reqStr)//拼接post提交body参数
	requestUrl := o.serverUrl + GetCorpAuthReqUrl              //接口请求api地址
	rspBody,requestId := common2.SendPost(requestUrl,headMap,[]byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody,&gerCorpAuthUrlRes)
	if err != nil{
		fmt.Println("json字符串转为struct失败")
	}
	gerCorpAuthUrlRes.RequestId = requestId
	return gerCorpAuthUrlRes
}


package client

import (
	"encoding/json"
	common2 "fasc_openapi_go_sdk/common"
	requestModel2 "fasc_openapi_go_sdk/model/requestModel"
	"fasc_openapi_go_sdk/model/responseModel"
	"fmt"
)

const(
	UserAddReqUrl ="/user/add" //添加个人用户
	GetUserAuthUrl ="/user/get-auth-url" //获取个人用户授权链接
)

// GetUserAddResponse 添加个人用户
func (o *openApiClient) GetUserAddResponse(userAddReq *requestModel2.UserAddReq,accessToken string) responseModel.UserAddRes  {
	var userAddRes responseModel.UserAddRes
	reqStr , err := json.Marshal(userAddReq)
	if err != nil{
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken,string(reqStr))
	bodyStr := "bizContent"+ "="+ string(reqStr) //拼接post提交body参数
	requestUrl := o.serverUrl + UserAddReqUrl              //接口请求api地址
	rspBody,requestId := common2.SendPost(requestUrl,headMap,[]byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody,&userAddRes)
	if err != nil{
		fmt.Println("json字符串转为struct失败")
	}
	userAddRes.RequestId =requestId
	return userAddRes
}

// GetUserAuthUrlResponse 获取个人用户收取按链接
func (o *openApiClient) GetUserAuthUrlResponse(authUrlReq *requestModel2.GetUserAuthUrlReq,accessToken string) responseModel.GetUserAuthUrlRes {
	var getAuthUrlRes responseModel.GetUserAuthUrlRes
	reqStr , err := json.Marshal(authUrlReq)
	if err != nil{
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken,string(reqStr))
	bodyStr := "bizContent"+ "="+ string(reqStr)//拼接post提交body参数
	requestUrl := o.serverUrl + GetUserAuthUrl              //接口请求api地址
	rspBody,requestId := common2.SendPost(requestUrl,headMap,[]byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody,&getAuthUrlRes)
	if err != nil{
		fmt.Println("json字符串转为struct失败")
	}
	getAuthUrlRes.RequestId = requestId
	return getAuthUrlRes
}




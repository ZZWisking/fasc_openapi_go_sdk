package client

import (
	"encoding/json"
	"fmt"
	common2 "github.com/ZZWisking/fasc_openapi_go_sdk/common"
	requestModel "github.com/ZZWisking/fasc_openapi_go_sdk/model/requestModel"
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/responseModel"
)

const (
	CorpAddReqUrl                = "/corp/add"                   //添加企业用户
	CorpUpdateReqUrl             = "/corp/update"                //更新企业用户
	CorpDeleteReqUrl             = "/corp/delete"                //删除企业用户
	CorpDisableReqUrl            = "/corp/disable"               //禁用企业用户
	CorpEnableReqUrl             = "/corp/enable"                //恢复企业用户
	GetCorpAuthReqUrl            = "/corp/get-auth-url"          //获取企业用户授权链接
	GetCorpInfoReqUrl            = "/corp/get"                   //查询企业用户基本信息
	GetCorpIdentityInfoReqUrl    = "/corp/get-identity-info"     //获取企业用户身份信息
	GetCorpIdentityProcessReqUrl = "/corp/get-identity-progress" //查询企业用户实名认证进度
)

// GetCorpAddResponse 添加企业用户
func (o *openApiClient) GetCorpAddResponse(corpAddReq *requestModel.CorpAddReq, accessToken string) responseModel.CorpAddRes {
	var corpAddRes responseModel.CorpAddRes
	reqStr, err := json.Marshal(corpAddReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + CorpAddReqUrl                                    //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &corpAddRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	corpAddRes.RequestId = requestId
	return corpAddRes
}

// GetCorpAuthUrlResponse 获取企业用户授权链接
func (o *openApiClient) GetCorpAuthUrlResponse(gerCorpAuthUrlReq *requestModel.GetCorpAuthUrlReq, accessToken string) responseModel.GetCorpAuthUrlRes {
	var gerCorpAuthUrlRes responseModel.GetCorpAuthUrlRes
	reqStr, err := json.Marshal(gerCorpAuthUrlReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetCorpAuthReqUrl                                //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &gerCorpAuthUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	gerCorpAuthUrlRes.RequestId = requestId
	return gerCorpAuthUrlRes
}

// GetCorpDeleteResponse 删除企业用户
func (o *openApiClient) GetCorpDeleteResponse(corpDeleteReq *requestModel.CorpDeleteReq, accessToken string) responseModel.CorpDeleteRes {
	var corpDeleteRes responseModel.CorpDeleteRes
	reqStr, err := json.Marshal(corpDeleteReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + CorpDeleteReqUrl                                 //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &corpDeleteRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	corpDeleteRes.RequestId = requestId
	return corpDeleteRes
}

// GetCorpUpdateResponse 更新企业用户
func (o *openApiClient) GetCorpUpdateResponse(corpUpdateReq *requestModel.CorpUpdateReq, accessToken string) responseModel.CorpUpdateRes {
	var corpUpdateRes responseModel.CorpUpdateRes
	reqStr, err := json.Marshal(corpUpdateReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + CorpUpdateReqUrl                                 //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &corpUpdateRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	corpUpdateRes.RequestId = requestId
	return corpUpdateRes
}

// GetCorpDisableResponse 禁用企业用户
func (o *openApiClient) GetCorpDisableResponse(corpDisableReq *requestModel.CorpDisableReq, accessToken string) responseModel.CorpDisableRes {
	var CorpDisableRes responseModel.CorpDisableRes
	reqStr, err := json.Marshal(corpDisableReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + CorpDisableReqUrl                                //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &CorpDisableRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	CorpDisableRes.RequestId = requestId
	return CorpDisableRes
}

// GetCorpEnableResponse 恢复企业用户
func (o *openApiClient) GetCorpEnableResponse(corpEnableReq *requestModel.CorpEnableReq, accessToken string) responseModel.CorpEnableRes {
	var corpEnableRes responseModel.CorpEnableRes
	reqStr, err := json.Marshal(corpEnableReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + CorpEnableReqUrl                                 //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &corpEnableRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	corpEnableRes.RequestId = requestId
	return corpEnableRes
}

// GetCorpInfoResponse 查询企业用户基本信息
func (o *openApiClient) GetCorpInfoResponse(corpInfoReq *requestModel.GetCorpReq, accessToken string) responseModel.GetCorpRes {
	var corpInfoRes responseModel.GetCorpRes
	reqStr, err := json.Marshal(corpInfoReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetCorpInfoReqUrl                                //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &corpInfoRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	corpInfoRes.RequestId = requestId
	return corpInfoRes
}

// GetCorpIdentityResponse 获取企业用户身份信息
func (o *openApiClient) GetCorpIdentityResponse(corpIdentityReq *requestModel.GetCorpIdentityInfoReq, accessToken string) responseModel.GetCorpIdentityInfoRes {
	var corpIdentityRes responseModel.GetCorpIdentityInfoRes
	reqStr, err := json.Marshal(corpIdentityReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetCorpIdentityInfoReqUrl                        //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &corpIdentityRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	corpIdentityRes.RequestId = requestId
	return corpIdentityRes
}

// GetCorpIdentityProcessResponse 查询企业用户实名认证进度
func (o *openApiClient) GetCorpIdentityProcessResponse(corpIdentityProcessReq *requestModel.GetCorpIdentityProgressReq, accessToken string) responseModel.GetCorpIdentityProgressRes {
	var corpIdentityProcessRes responseModel.GetCorpIdentityProgressRes
	reqStr, err := json.Marshal(corpIdentityProcessReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetCorpIdentityProcessReqUrl                     //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &corpIdentityProcessRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	corpIdentityProcessRes.RequestId = requestId
	return corpIdentityProcessRes
}

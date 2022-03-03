package client

import (
	"encoding/json"
	"fmt"
	common2 "github.com/ZZWisking/fasc_openapi_go_sdk/common"
	requestModel "github.com/ZZWisking/fasc_openapi_go_sdk/model/requestModel"
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/responseModel"
)

const (
	UserAddReqUrl            = "/user/add"                   //添加个人用户
	UserUpdateReqUrl         = "/user/update"                //更新个人用户
	UserDeleteReqUrl         = "/user/delete"                //删除个人用户
	UserDisableReqUrl        = "/user/disable"               //禁用个人用户
	UserEnableReqUrl         = "/user/enable"                //恢复个人用户
	GetUserAuthUrl           = "/user/get-auth-url"          //获取个人用户授权链接
	GetUserInfoReqUrl        = "/user/get"                   //查询个人用户基本信息
	GetIdentityInfoReqUrl    = "/user/get-identity-info"     //获取个人用户身份信息
	GetIdentityProcessReqUrl = "/user/get-identity-progress" //查询个人用户实名认证进度
)

// GetUserAddResponse 添加个人用户
func (o *openApiClient) GetUserAddResponse(userAddReq *requestModel.UserAddReq, accessToken string) responseModel.UserAddRes {
	var userAddRes responseModel.UserAddRes
	reqStr, err := json.Marshal(userAddReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + UserAddReqUrl                                    //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &userAddRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userAddRes.RequestId = requestId
	return userAddRes
}

// GetUserAuthUrlResponse 获取个人用户收取按链接
func (o *openApiClient) GetUserAuthUrlResponse(authUrlReq *requestModel.GetUserAuthUrlReq, accessToken string) responseModel.GetUserAuthUrlRes {
	var getAuthUrlRes responseModel.GetUserAuthUrlRes
	reqStr, err := json.Marshal(authUrlReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetUserAuthUrl                                   //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &getAuthUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	getAuthUrlRes.RequestId = requestId
	return getAuthUrlRes
}

// GetUserDeleteResponse 删除个人用户
func (o *openApiClient) GetUserDeleteResponse(userDeleteReq *requestModel.UserDeleteReq, accessToken string) responseModel.UserDeleteRes {
	var userDeleteRes responseModel.UserDeleteRes
	reqStr, err := json.Marshal(userDeleteReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + UserDeleteReqUrl                                 //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &userDeleteRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userDeleteRes.RequestId = requestId
	return userDeleteRes
}

// GetUserUpdateResponse 更新个人用户
func (o *openApiClient) GetUserUpdateResponse(userUpdateReq *requestModel.UserUpdateReq, accessToken string) responseModel.UserUpdateRes {
	var userUpdateRes responseModel.UserUpdateRes
	reqStr, err := json.Marshal(userUpdateReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + UserUpdateReqUrl                                 //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &userUpdateRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userUpdateRes.RequestId = requestId
	return userUpdateRes
}

// GetUserDisableResponse 禁用个人用户
func (o *openApiClient) GetUserDisableResponse(userDisableReq *requestModel.UserDisableReq, accessToken string) responseModel.UserDisableRes {
	var userDisableRes responseModel.UserDisableRes
	reqStr, err := json.Marshal(userDisableReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + UserDisableReqUrl                                //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &userDisableRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userDisableRes.RequestId = requestId
	return userDisableRes
}

// GetUserEnableResponse 恢复个人用户
func (o *openApiClient) GetUserEnableResponse(userEnableReq *requestModel.UserEnableReq, accessToken string) responseModel.UserEnableRes {
	var userEnableRes responseModel.UserEnableRes
	reqStr, err := json.Marshal(userEnableReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + UserEnableReqUrl                                 //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &userEnableRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userEnableRes.RequestId = requestId
	return userEnableRes
}

// GetUserInfoResponse 查询个人用户基本信息
func (o *openApiClient) GetUserInfoResponse(userInfoReq *requestModel.GetUserReq, accessToken string) responseModel.GetUserRes {
	var userInfoRes responseModel.GetUserRes
	reqStr, err := json.Marshal(userInfoReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetUserInfoReqUrl                                //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &userInfoRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userInfoRes.RequestId = requestId
	return userInfoRes
}

// GetUserIdentityResponse 获取个人用户身份信息
func (o *openApiClient) GetUserIdentityResponse(userIdentityReq *requestModel.GetUserIdentifyInfoReq, accessToken string) responseModel.GetUserIdentifyInfoRes {
	var userIdentityRes responseModel.GetUserIdentifyInfoRes
	reqStr, err := json.Marshal(userIdentityReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetIdentityInfoReqUrl                            //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &userIdentityRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userIdentityRes.RequestId = requestId
	return userIdentityRes
}

// GetUserIdentityProcessResponse 查询个人用户实名认证进度
func (o *openApiClient) GetUserIdentityProcessResponse(userIdentityProcessReq *requestModel.GetUserIdentityProgressReq, accessToken string) responseModel.GetUserIdentityProgressRes {
	var userIdentityProcessRes responseModel.GetUserIdentityProgressRes
	reqStr, err := json.Marshal(userIdentityProcessReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetIdentityProcessReqUrl                         //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &userIdentityProcessRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	userIdentityProcessRes.RequestId = requestId
	return userIdentityProcessRes
}

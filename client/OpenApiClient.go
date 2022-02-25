package client

import (
	"encoding/json"
	common2 "fasc_openapi_go_sdk/common"
	"fasc_openapi_go_sdk/model"
	"fmt"
	"strconv"
	"time"
)

type openApiClient struct {
	appId string
	appSecret string
	serverUrl string
}
const (
	signType ="HMAC-SHA256" //加密算法
	reqUrl = "/service/get-access-token" //获取服务访问凭证
	resCode = "100000" //接口请求成功状态码
)

func NewClient(AppID, AppSecret string,ServerUrl string) *openApiClient {
	return &openApiClient{
		appId: AppID,
		appSecret: AppSecret,
		serverUrl: ServerUrl,
	}
}


// GetAuthToken 获取服务凭证
func(o *openApiClient) GetAuthToken() model.AccessTokenRes {
	  var accessTokenRes model.AccessTokenRes
      requestUrl := o.serverUrl + reqUrl                               //接口请求api地址
      rspBody,_ := common2.SendPost(requestUrl,o.GetReqHeadParams(),nil) //发送post请求
	  err := json.Unmarshal(rspBody,&accessTokenRes)
	  if err != nil{
		  fmt.Println("json字符串转为struct失败")
	  }
      return accessTokenRes
}

// GetReqHeadParams 拼接获取服务凭证请求参数
func(o *openApiClient) GetReqHeadParams() map[string]string {
	var headMap map[string]string
	headMap = make(map[string]string)
	nonce := common2.GetGUID().Hex()
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	headMap["X-FASC-App-Id"] =o.appId
	headMap["X-FASC-Sign-Type"]= signType
	headMap["X-FASC-Timestamp"] =timestamp
	headMap["X-FASC-Nonce"] =nonce
	headMap["X-FASC-Grant-Type"] ="client_credential"
	signature := common2.GetSignByMap(headMap,timestamp,o.appSecret)
	headMap["X-FASC-Sign"] =signature
	return headMap
}

// SetReqHeadMap 拼接包含请求参数的head请求参数
func (o *openApiClient) SetReqHeadMap(accessToken string,bizContent string) map[string]string {
	var headMap map[string]string
	headMap = make(map[string]string)
	nonce := common2.GetGUID().Hex()
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	headMap["X-FASC-App-Id"] =o.appId
	headMap["X-FASC-Sign-Type"]= signType
	headMap["X-FASC-Timestamp"] =timestamp
	headMap["X-FASC-Nonce"] =nonce
	headMap["X-FASC-AccessToken"] = accessToken
	headMap["bizContent"] = bizContent
	signature := common2.GetSignByMap(headMap,timestamp,o.appSecret)
	headMap["X-FASC-Sign"] =signature
	delete(headMap, "bizContent")
	return headMap
}



package client

import (
	"encoding/json"
	common2 "github.com/ZZWisking/fasc_openapi_go_sdk/common"
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/requestModel"
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/responseModel"
	"fmt"
)

const (
	GetManageUrlReqUrl = "/manage/get-manage-url" //获取企业管理链接
)

// GetCorpManageUrlResponse 获取企业管理链接
func (o *openApiClient) GetCorpManageUrlResponse(corpManageUrlReq *requestModel.GetCorpManageUrlReq, accessToken string) responseModel.GetCorpManageUrlRes {
	var corpManageUrlRes responseModel.GetCorpManageUrlRes
	reqStr, err := json.Marshal(corpManageUrlReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetManageUrlReqUrl                               //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &corpManageUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	corpManageUrlRes.RequestId = requestId
	return corpManageUrlRes
}

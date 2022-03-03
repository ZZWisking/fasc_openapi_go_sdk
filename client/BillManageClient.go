package client

import (
	"encoding/json"
	"fmt"
	common2 "github.com/ZZWisking/fasc_openapi_go_sdk/common"
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/requestModel"
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/responseModel"
)

const (
	GetBillUrlReqUrl = "/billing/get-bill-url" //获取计费链接
)

// GetBillUrlResponse 获取计费链接
func (o *openApiClient) GetBillUrlResponse(billUrlReq *requestModel.GetBillingUrlReq, accessToken string) responseModel.GetBillingUrlRes {
	var billUrlRes responseModel.GetBillingUrlRes
	reqStr, err := json.Marshal(billUrlReq)
	if err != nil {
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken, string(reqStr))
	bodyStr := "bizContent" + "=" + string(reqStr)                               //拼接post提交body参数
	requestUrl := o.serverUrl + GetBillUrlReqUrl                                 //接口请求api地址
	rspBody, requestId := common2.SendPost(requestUrl, headMap, []byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody, &billUrlRes)
	if err != nil {
		fmt.Println("json字符串转为struct失败")
	}
	billUrlRes.RequestId = requestId
	return billUrlRes
}

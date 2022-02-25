package client

import (
	"encoding/json"
	"fmt"
	common2 "github.com/ZZWisking/fasc_openapi_go_sdk/common"
	requestModel2 "github.com/ZZWisking/fasc_openapi_go_sdk/model/requestModel"
	"github.com/ZZWisking/fasc_openapi_go_sdk/model/responseModel"
)

const(
	GetSignTemplateListReqUrl ="/sign-template/get-list" //获取签署模板列表
)

// GetSignTemplateListResponse 查询签署模板列表
func (o *openApiClient) GetSignTemplateListResponse(signTemplateReq *requestModel2.GetSignTemplateListReq,accessToken string) responseModel.GetSignTemplateListRes  {
	var signTemplateRes responseModel.GetSignTemplateListRes
	reqStr , err := json.Marshal(signTemplateReq)
	if err != nil{
		fmt.Println("json序列化失败")
	}
	headMap := o.SetReqHeadMap(accessToken,string(reqStr))
	bodyStr := "bizContent"+ "="+ string(reqStr)//拼接post提交body参数
	requestUrl := o.serverUrl + GetSignTemplateListReqUrl         //接口请求api地址
	rspBody,requestId := common2.SendPost(requestUrl,headMap,[]byte(bodyStr)) //发送post请求
	err = json.Unmarshal(rspBody,&signTemplateRes)
	if err != nil{
		fmt.Println("json字符串转为struct失败")
	}
	signTemplateRes.RequestId = requestId
	return signTemplateRes
}


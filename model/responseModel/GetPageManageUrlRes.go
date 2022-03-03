package responseModel

type PageManageUrlResData struct {
	PUrl string `json:"pUrl"`
}

// GetPageManageUrlRes 获取模板管理链接
type GetPageManageUrlRes struct {
	RequestId string               `json:"requestId"`
	Code      string               `json:"code"`
	Msg       string               `json:"msg"`
	Data      PageManageUrlResData `json:"data"`
}

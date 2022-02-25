package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// SendPost httpPost 请求
func SendPost(urlStr string,headMaps map[string]string,data []byte) ([]byte,string) {
	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewReader(data)) // URL-encoded payload
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headMaps{
		r.Header.Add(k, v)
	}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return nil,""
	}
	defer resp.Body.Close()
	requestId := resp.Header.Values("X-FASC-Request-Id")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
	}
	if requestId == nil{
		return body,""
	}
	return body,requestId[0]
}

// DownLoadFilesPost 下载文件HttpPost
func DownLoadFilesPost(urlStr string,headMaps map[string]string,data []byte) ([]byte,string,string) {
	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewReader(data)) // URL-encoded payload
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headMaps{
		r.Header.Add(k, v)
	}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return nil,"",""
	}
	defer resp.Body.Close()
	requestId := resp.Header.Values("X-FASC-Request-Id")
	contentType := resp.Header.Values("Content-Type")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
	}
	if requestId == nil{
		return body,"",contentType[0]
	}
	if contentType == nil{
		return body,requestId[0],""
	}
	return body,requestId[0],contentType[0]
}

func PostWithFormData(method, url string, postData *map[string]string){
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k,v :=  range *postData{
		w.WriteField(k, v)
	}
	w.Close()
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, _ := http.DefaultClient.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Printf("%s", data)
}


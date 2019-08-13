package utils

import (
	"github.com/kataras/iris"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 提交get请求
func ReqGet(urlPath string, timeout time.Duration) []byte {
	result := []byte("") //默认的返回结果
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", urlPath, strings.NewReader(""))
	if err != nil {
		return result
	}
	resp, err := client.Do(req)
	if err != nil {
		return result
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result
	}
	return body
}

// 发送post请求
func ReqPost(urlPath string, data map[string]string, timeout time.Duration) []byte {
	result := []byte("{}")
	client := &http.Client{
		Timeout: timeout,
	}
	httpBuildQuery := ""
	for k, v := range data {
		//如果传进来的是已经拼接好的，就放入map,k的值就是拼接好的,value为空字符串
		if len(data) == 1 && v == "" {
			httpBuildQuery = k
		} else {
			httpBuildQuery += k + "=" + v + "&"
		}
	}
	if httpBuildQuery != "" {
		httpBuildQuery = strings.TrimRight(httpBuildQuery, "&")
	}
	req, err := http.NewRequest("POST", urlPath, strings.NewReader(httpBuildQuery))
	if err != nil {
		return result
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return result
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("{}")
	}
	return body
}



func CurlPostJson(urlPath string, data string, timeout time.Duration) []byte {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", urlPath, strings.NewReader(data))
	if err != nil {
		return []byte("{}")
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []byte("{}")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}
//得到提交的form表单当中的数据
//解析传过来的post数据, 因为是 body形式整体传过来的json数据，故需要对期进行解析
//如果前端代码有变动，此处需要做出相应的调整
func GetPostData(ctx *iris.Context) PostData {

	contentType := (*ctx).GetHeader("content-type")
	postType := 0 //默认是form
	if strings.Index(contentType, "/json") > 0 {
		postType = 1
	}

	if postType == 0 {
		return PostData{
			Type: postType,
			Ctx:  ctx,
			Data: &[]byte{},
		}
	}

	body := (*ctx).Params().Get("PostBody")
	if strings.Compare(body, "") == 0 {
		data, _ := ioutil.ReadAll((*ctx).Request().Body)
		postBody := string(data)
		(*ctx).Params().SetImmutable("PostBody", postBody)
		return PostData{
			Type: postType,
			Ctx:  ctx,
			Data: &data,
		}
	}

	bytes := []byte(body)
	return PostData{
		Type: postType,
		Ctx:  ctx,
		Data: &bytes,
	}
}

//路径拼接,http表单处理
func UrlSplitKeyValueOnlyHttp(urlPath string, data map[string]string, isEncode bool, isHttp bool) string {

	if isHttp {
		urlPath = strings.Replace(urlPath, "https://", "http://", 1)
	}
	urlPath += "?"
	for k, v := range data {
		if isEncode {
			v = url.QueryEscape(v)
		}
		urlPath += k + "=" + v + "&"
	}
	urlPath = strings.TrimRight(urlPath, "&")

	return urlPath
}

//路径拼接
func UrlSplitKeyValue(urlPath string, data map[string]string, isEncode bool) string {
	urlPath += "?"
	for k, v := range data {
		if isEncode {
			v = url.QueryEscape(v)
		}
		urlPath += k + "=" + v + "&"
	}
	urlPath = strings.TrimRight(urlPath, "&")
	return urlPath
}

func BuildUrl(urlPath string, params map[string]string) string {
	var fullUrl string
	uv := url.Values{}
	for k, v := range params {
		uv.Add(k, v)
	}
	newQueries := uv.Encode()
	if strings.Contains(urlPath, "?") {
		fullUrl = urlPath + "&" + newQueries
	} else {
		fullUrl = urlPath + "?" + newQueries
	}
	return fullUrl
}

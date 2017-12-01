package cnet

import (
	"net/http"
	"io/ioutil"
	"strings"
)

/**
 * 发起HTTP请求
 * param string url: 请求地址
 * param string ref: http header Referer信息
 * param string cookie: http header cookie
 * param string auth: http header Authorization
 * param string data: http header data
 */
func HttpRequest(url string, ref string, cookie string, auth string, data string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, strings.NewReader(data))
	if err != nil {
		// handle error
	}

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; rv:37.0) Gecko/20100101 Firefox/37.0")

	if len(cookie) > 0 {
		req.Header.Set("Cookie", cookie)
	}

	if len(auth) > 0 {
		req.Header.Set("Authorization", "Bearer " + auth)
	}

	if len(ref) > 0 {
		req.Header.Set("Referer", ref)
	}

	// 句柄不关闭会造成内存泄露
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body
}

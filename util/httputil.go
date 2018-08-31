package util

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//HttpClient http请求
type HttpClient struct {
	TimeOut time.Duration //设置超时时间默认10s
	client  *http.Client
	Log     *log.Logger //日志记录
}

//NewHttpClient 新建一个请求
func NewHttpClient(timeout time.Duration) *HttpClient {
	if timeout <= 0 {
		timeout = time.Second * 10
	}
	cli := &http.Client{Transport: http.DefaultTransport}
	log := log.New(os.Stdout, "OpenApi_Test  ", log.Lshortfile|log.Ldate|log.Lmicroseconds|log.LstdFlags)
	return &HttpClient{
		TimeOut: timeout,
		client:  cli,
		Log:     log,
	}
}

//UseProxy 设置代理
func (cli *HttpClient) UseProxy(proxyurl *url.URL) *HttpClient {
	proxy := func(req *http.Request) (*url.URL, error) {
		return proxyurl, nil
	}
	trans := &http.Transport{Proxy: proxy}
	cli.client.Transport = trans
	return cli

}

//HTTPPost 请求
func (cli *HttpClient) HTTPPost(url string, header map[string]string, body map[string]interface{}) (string, error) {
	var reader io.Reader
	if body != nil && len(body) > 0 {
		params := make([]string, 0)
		for k, v := range body {
			params = append(params, fmt.Sprintf("%s=%v", k, v))
		}
		reader = strings.NewReader(strings.Join(params, "&"))
	}

	req, err := http.NewRequest(http.MethodPost, url, reader)

	if err != nil {
		return "", err
	}
	if header != nil && len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	bodycopy, _ := req.GetBody()
	content, _ := ioutil.ReadAll(bodycopy)
	bodycopy.Close()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	cli.Log.Printf("RequestUrl:%s;Method:%s;TimeOut:%v; Header:%v;Body:%v", req.URL, req.Method, cli.client.Timeout, req.Header, string(content))
	resp, err2 := cli.client.Do(req)
	if err2 != nil {
		return "", err2
	}
	defer resp.Body.Close()
	respbodys, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return "", err3
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 600 {
		return "", errors.New(string(respbodys))
	}
	return string(respbodys), err3
}

//HTTPGet 请求
func (cli *HttpClient) HTTPGet(url string, header map[string]string, param map[string]interface{}) (string, error) {
	var reader io.Reader

	if param != nil && len(param) > 0 {
		params := make([]string, 0)
		for k, v := range param {
			params = append(params, fmt.Sprintf("%s=%v", k, v))
		}
		reader = strings.NewReader(strings.Join(params, "&"))
		query, _ := ioutil.ReadAll(reader)
		if strings.Contains(url, "?") {
			url = url + "&" + string(query)
		} else {
			url = url + "?" + string(query)
		}
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}

	if header != nil && len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	client := &http.Client{Timeout: cli.TimeOut}
	cli.Log.Printf("RequestUrl:%s;Method:%s;TimeOut:%v; Header:%v", req.URL, req.Method, cli.client.Timeout, req.Header)
	resp, err2 := client.Do(req)
	if err2 != nil {
		return "", err2
	}
	defer resp.Body.Close()
	respbodys, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return "", err3
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 600 {
		return "", errors.New(string(respbodys))
	}
	return string(respbodys), err3
}

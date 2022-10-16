package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-admin/app/patent/my_config"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	authUrl   = "http://www.innojoy.com/accountAuth.aspx"
	searchUrl = "http://www.innojoy.com/service/patentSearch.aspx"
)

type httpClient struct {
	client *http.Client
}

func newHttpClient() *httpClient {
	httpTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时时间
			KeepAlive: 60 * time.Second, // 保持长连接的时间
		}).DialContext, // 设置连接的参数
		MaxIdleConns:          500,              // 最大空闲连接
		IdleConnTimeout:       60 * time.Second, // 空闲连接的超时时间
		ExpectContinueTimeout: 30 * time.Second, // 等待服务第一个响应的超时时间
		MaxIdleConnsPerHost:   100,              // 每个host保持的空闲连接数
	}

	return &httpClient{
		client: &http.Client{
			Transport: httpTransport,
		}}
}

func (hc *httpClient) Get(uri string, params map[string]string) (*http.Response, error) {
	return hc.requestWithoutBody(uri, params, nil, http.MethodGet)
}

//func (hc *httpClient) GetAndRead(uri string, params map[string]string, successCode int) ([]byte, error) {
//	resp, err := hc.Get(uri, params)
//	if err != nil {
//		return nil, errno.New(errno.InternalServerError, err)
//	}
//	buf, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, errno.New(errno.InternalServerError, err)
//	}
//	if resp.StatusCode != successCode {
//		erro := &errno.Errno{
//			Code:    resp.StatusCode,
//			Message: string(buf),
//		}
//		return nil, errno.New(erro, fmt.Errorf("invalid status code: %d", resp.StatusCode))
//	}
//	return buf, nil
//}
//
//func (hc *httpClient) GetAndUnmarshal(uri string, params map[string]string, successCode int, obj interface{}) error {
//	resp, err := hc.Get(uri, params)
//	if err != nil {
//		return errno.New(errno.InternalServerError, err)
//	}
//	buf, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return errno.New(errno.InternalServerError, err)
//	}
//	if resp.StatusCode != successCode {
//		erro := &errno.Errno{
//			Code:    resp.StatusCode,
//			Message: string(buf),
//		}
//		return errno.New(erro, fmt.Errorf("invalid status code: %d", resp.StatusCode))
//	}
//	err = json.Unmarshal(buf, obj)
//	if err != nil {
//		return errno.New(errno.InternalServerError, err)
//	}
//	return err
//}

func (hc *httpClient) Post(url string, body interface{}, params map[string]string) (*http.Response, error) {
	return hc.requestWithBody(url, body, params, nil, http.MethodPost)
}

func (hc *httpClient) Put(url string, body interface{}, params map[string]string) (*http.Response, error) {
	return hc.requestWithBody(url, body, params, nil, http.MethodPut)
}

func (hc *httpClient) Delete(uri string, params map[string]string) (*http.Response, error) {
	return hc.requestWithoutBody(uri, params, nil, http.MethodDelete)
}

func (hc *httpClient) requestWithoutBody(uri string, params map[string]string, headers map[string]string, method string) (*http.Response, error) {
	req, err := http.NewRequest(method, uri, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	q := req.URL.Query()

	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	// add auth for pulsar_admin
	//req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", hc.token))

	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	return hc.client.Do(req)
}

func (hc *httpClient) requestWithBody(url string, body interface{}, params map[string]string, headers map[string]string, method string) (*http.Response, error) {
	var bodyJSON []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJSON, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return nil, errors.New("http post body to json failed")
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(fmt.Sprintf("new request is fail: %v \n", err))
	}

	if method == http.MethodPatch {
		req.Header.Set("Content-type", "application/json-patch+json")
	} else {
		req.Header.Set("Content-type", "application/json;charset=utf-8")
	}

	q := req.URL.Query()

	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	// add auth for pulsar_admin
	//req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", hc.token))

	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	return hc.client.Do(req)
}

type loginReq struct {
	UserConfig UserConfig `json:"userConfig"`
}

type UserConfig struct {
	EMail    string `json:"EMail"`
	Password string `json:"Password"`
}

type loginResp struct {
	ReturnValue int `json:"ReturnValue"`
	Option      struct {
		Token string `json:"token"`
	} `json:"Option"`
	ErrorInfo string `json:"ErrorInfo"`
}

type innojoyClient struct {
	hc    *httpClient
	token string
}

func newInnojoyClient() *innojoyClient {
	return &innojoyClient{
		hc: newHttpClient(),
	}
}

func (ic *innojoyClient) autoLogin() (string, error) {
	req := &loginReq{UserConfig: UserConfig{
		EMail:    my_config.CurrentPatentConfig.InnojoyUser,
		Password: my_config.CurrentPatentConfig.InnojoyPassword,
	}}

	resp, err := ic.hc.Post(authUrl, req, nil)
	if err != nil {
		return "", err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	loginRes := loginResp{}
	if err = json.Unmarshal(buf, &loginRes); err != nil {
		return "", err
	}

	return loginRes.Option.Token, nil
}

func (ic *innojoyClient) simpleSearch() {

}

//func (ic *innojoyClient) search(sr *searchReq) {
//	resp, err :=ic.hc.Post(searchUrl, sr, nil)
//
//}

type searchReq struct {
	Token              string `json:"token"`
	PatentSearchConfig struct {
		GUID      string `json:"GUID"`
		Action    string `json:"Action"`
		Query     string `json:"Query"`
		Database  string `json:"Database"`
		Page      string `json:"Page"`
		PageSize  string `json:"PageSize"`
		Sortby    string `json:"Sortby"`
		FieldList string `json:"FieldList"`
	} `json:"patentSearchConfig"`
}

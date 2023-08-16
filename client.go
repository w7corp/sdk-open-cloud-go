package w7

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/w7corp/sdk-open-cloud-go/service"
	"log"
	"net/url"
	"sort"
)

type Option struct {
	ApiUrl         string
	Debug          bool
	DefaultHeaders map[string]string
}

func NewClient(appId string, appSecret string, option Option) *Client {
	client := &Client{
		appId:     appId,
		appSecret: appSecret,
	}
	client.apiUrl = "https://api.w7.cc"
	if option.ApiUrl != "" {
		client.apiUrl = option.ApiUrl
	}

	httpClient := resty.New()
	httpClient.SetBaseURL(client.apiUrl)
	httpClient.OnBeforeRequest(client.makeSign)
	if option.Debug {
		httpClient.EnableTrace()
	}
	if option.DefaultHeaders != nil {
		httpClient.SetHeaders(option.DefaultHeaders)
	}
	client.httpClient = httpClient

	client.OauthService = &service.OauthService{
		HttpClient: httpClient,
	}

	return client
}

type Client struct {
	apiUrl    string
	appId     string
	appSecret string

	httpClient *resty.Client

	OauthService *service.OauthService
}

func (c *Client) GetHttpClient() *resty.Client {
	return c.httpClient
}

func (c *Client) makeSign(client *resty.Client, request *resty.Request) error {
	request.SetFormData(map[string]string{
		"appid": c.appId,
	})
	var keys []string
	var signStr string
	for s, _ := range request.FormData {
		if s == "sign" {
			continue
		}
		keys = append(keys, s)
	}
	sort.Strings(keys)
	for i, k := range keys {
		signStr += fmt.Sprintf("%s=%s", k, url.QueryEscape(request.FormData.Get(k)))
		if i < len(keys)-1 {
			signStr += "&"
		}
	}
	signStr += c.appSecret
	log.Printf("签名数据：%s \n", signStr)

	sign := md5.Sum([]byte(signStr))
	request.SetFormData(map[string]string{
		"sign": hex.EncodeToString(sign[:]),
	})
	return nil
}

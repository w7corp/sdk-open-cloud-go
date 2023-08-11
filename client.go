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
	ApiUrl string
}

func NewClient(appId string, appSecret string, options ...Option) *Client {
	client := &Client{
		appId:     appId,
		appSecret: appSecret,
	}

	client.apiUrl = "https://api.w7.cc"

	for _, option := range options {
		if option.ApiUrl != "" {
			client.apiUrl = option.ApiUrl
		}
	}

	httpClient := resty.New()
	httpClient.SetBaseURL(client.apiUrl)
	httpClient.OnBeforeRequest(client.makeSign)
	client.OauthService = &service.OauthService{
		HttpClient: httpClient,
	}

	return client
}

type Client struct {
	apiUrl    string
	appId     string
	appSecret string

	OauthService *service.OauthService
}

func (self *Client) makeSign(client *resty.Client, request *resty.Request) error {
	request.SetFormData(map[string]string{
		"appid": self.appId,
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
	signStr += self.appSecret
	log.Printf("签名数据：%s \n", signStr)

	sign := md5.Sum([]byte(signStr))
	request.SetFormData(map[string]string{
		"sign": hex.EncodeToString(sign[:]),
	})
	return nil
}

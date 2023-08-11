package w7

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/w7corp/sdk-open-cloud-go/service"
	"log"
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type Option struct {
	ApiUrl string
	Debug  bool
}

func NewClient(appId string, appSecret string, options ...Option) *Client {
	client := &Client{
		appId:     appId,
		appSecret: appSecret,
	}
	client.log = &wlog{}
	client.apiUrl = "https://api.w7.cc"

	for _, option := range options {
		if option.ApiUrl != "" {
			client.apiUrl = option.ApiUrl
		}
		if option.Debug {
			client.log.debug = option.Debug
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
	apiUrl       string
	appId        string
	appSecret    string
	log          *wlog
	OauthService *service.OauthService
}

func (self *Client) makeSign(client *resty.Client, request *resty.Request) error {
	request.SetFormData(map[string]string{
		"appid":     self.appId,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
		"nonce":     self.getRandomString(5),
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
	self.log.Printf("签名数据：%s \n", signStr)

	sign := md5.Sum([]byte(signStr))
	request.SetFormData(map[string]string{
		"sign": hex.EncodeToString(sign[:]),
	})
	self.log.Printf("签名：%s \n", hex.EncodeToString(sign[:]))
	return nil
}

func (self *Client) getRandomString(n int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

type wlog struct {
	debug bool
}

func (self *wlog) Println(v ...any) {
	if self.debug {
		log.Println(v...)
	}
}

func (self *wlog) Printf(format string, v ...any) {
	if self.debug {
		log.Printf(format, v...)
	}
}

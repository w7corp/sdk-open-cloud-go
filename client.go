package w7

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/w7corp/sdk-open-cloud-go/service"
	"golang.org/x/net/publicsuffix"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/http/cookiejar"
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

	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	httpClient := resty.NewWithClient(&http.Client{
		Jar: cookieJar,
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           dialer.DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          1000,
			MaxIdleConnsPerHost:   math.MaxInt,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	})
	httpClient.SetBaseURL(client.apiUrl)
	httpClient.OnBeforeRequest(client.makeSign)
	httpClient.OnAfterResponse(client.onafterResponse)
	client.SetHttpClient(httpClient)

	client.OauthService = &service.OauthService{
		HttpClient: httpClient,
	}

	return client
}

type Client struct {
	apiUrl     string
	appId      string
	appSecret  string
	log        *wlog
	httpClient *resty.Client

	OauthService *service.OauthService
}

func (c *Client) SetHttpClient(client *resty.Client) {
	c.httpClient = client
}

func (c *Client) GetHttpClient() *resty.Client {
	return c.httpClient
}

func (c *Client) makeSign(client *resty.Client, request *resty.Request) error {
	var sign [16]byte
	if request.Body != nil && resty.DetectContentType(request.Body) == "application/json" {
		body, ok := request.Body.(map[string]interface{})
		if !ok {
			return errors.New("request property body must be ")
		}
		body["appid"] = c.appId
		body["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
		body["nonce"] = c.getRandomString(16)

		signByte, err := client.JSONMarshal(body)
		if err != nil {
			return err
		}
		signStr := string(signByte)
		signStr += c.appSecret
		sign = md5.Sum([]byte(signStr))

		body["sign"] = hex.EncodeToString(sign[:])
		request.SetBody(body)
	} else {
		request.SetFormData(map[string]string{
			"appid":     c.appId,
			"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
			"nonce":     c.getRandomString(16),
		})
		var keys []string
		signStr := ""
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
		c.log.Printf("签名数据：%s \n", signStr)

		sign = md5.Sum([]byte(signStr))
		request.SetFormData(map[string]string{
			"sign": hex.EncodeToString(sign[:]),
		})
	}
	c.log.Printf("签名：%s \n", hex.EncodeToString(sign[:]))
	return nil
}

func (c *Client) onafterResponse(client *resty.Client, response *resty.Response) error {
	c.log.Println("response data: " + string(response.Body()))
	return nil
}

func (c *Client) getRandomString(n int) string {
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

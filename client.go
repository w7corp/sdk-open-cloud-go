package w7

import (
	"github.com/go-resty/resty/v2"
	"github.com/w7corp/sdk-open-cloud-go/service"
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
	return nil
}

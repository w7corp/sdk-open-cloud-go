package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	w7 "github.com/w7corp/sdk-open-cloud-go"
	"testing"
)

func TestOauth(t *testing.T) {
	client := w7.NewClient(APP_ID, APP_SECRET, w7.Option{
		Debug: true,
	})
	loginUrl, err := client.OauthService.GetLoginUrl("http://s.w7.cc")

	if err != nil {
		t.Errorf("message: %s, code: %d", err.Message, err.Errno)
		return
	}
	test := assert.New(t)
	test.NotEmpty(loginUrl)

	fmt.Println(loginUrl)
}

func TestAccessToken(t *testing.T) {
	client := w7.NewClient(APP_ID, APP_SECRET)
	accessToken, err := client.OauthService.GetAccessTokenByCode("123456789")

	if err != nil {
		if err.Message == "Code is already in use" {
			t.Log(err.Message)
			return
		}
		t.Errorf("message: %s, code: %d", err.Message, err.Errno)
		return
	}
	test := assert.New(t)
	test.NotEmpty(accessToken.AccessToken)
	println(accessToken.AccessToken)
}

func TestUserInfo(t *testing.T) {
	client := w7.NewClient(APP_ID, APP_SECRET)
	userInfo, err := client.OauthService.GetUserInfo(ACCESS_TOKEN)

	if err != nil {
		t.Errorf("message: %s, code: %d", err.Message, err.Errno)
		return
	}
	test := assert.New(t)
	test.NotEmpty(userInfo.OpenId)
	fmt.Printf("result user openid: %s \n", userInfo.OpenId)
}

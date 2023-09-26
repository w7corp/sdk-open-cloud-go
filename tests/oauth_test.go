package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	w7 "github.com/w7corp/sdk-open-cloud-go"
	"testing"
)

var APP_ID = ""
var APP_SECRET = ""
var ACCESS_TOKEN = ""

func TestOauth(t *testing.T) {
	client := w7.NewClient(APP_ID, APP_SECRET, w7.Option{
		Debug: true,
	})
	loginUrl, err := client.OauthService.GetLoginUrl("http://s.w7.cc")

	if err.IsError() {
		t.Errorf("message: %s, code: %d", err.ErrMsg, err.Errno)
		return
	}
	test := assert.New(t)
	test.NotEmpty(loginUrl)

	fmt.Println(loginUrl)
}

func TestAccessToken(t *testing.T) {
	client := w7.NewClient(APP_ID, APP_SECRET, w7.Option{
		Debug: true,
	})
	accessToken, err := client.OauthService.GetAccessTokenByCode("123456789")

	if err.IsError() {
		if err.ErrMsg == "Code is already in use" {
			t.Log(err.ErrMsg)
			return
		}
		t.Errorf("message: %s, code: %d", err.ErrMsg, err.Errno)
		return
	}
	test := assert.New(t)
	test.NotEmpty(accessToken.AccessToken)
	println(accessToken.AccessToken)
}

func TestUserInfo(t *testing.T) {
	client := w7.NewClient(APP_ID, APP_SECRET, w7.Option{
		Debug: true,
	})
	userInfo, err := client.OauthService.GetUserInfo(ACCESS_TOKEN)

	if err.IsError() {
		t.Errorf("message: %s, code: %d", err.ErrMsg, err.Errno)
		return
	}
	test := assert.New(t)
	test.NotEmpty(userInfo.OpenId)
	fmt.Printf("result user openid: %s \n", userInfo.OpenId)
}

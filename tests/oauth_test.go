package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	w7 "github.com/w7corp/sdk-open-cloud-go"
	"testing"
)

func TestOauth(t *testing.T) {
	client := w7.NewClient(APP_ID, APP_SECRET)
	loginUrl, err := client.OauthService.GetLoginUrl("http://s.w7.cc")

	if err != nil {
		t.Errorf("message: %s, code: %d", err.Message, err.Errno)
		return
	}
	test := assert.New(t)
	test.NotEmpty(loginUrl)

	fmt.Println(loginUrl)
}

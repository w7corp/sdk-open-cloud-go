package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	w7 "github.com/w7corp/sdk-open-cloud-go"
	"testing"
)

func TestSign(t *testing.T) {
	client := w7.NewClient("wa84a4166e8e1f471a", "4qj8BJtSS0F9qM2dCHog/NT9BRzxJtOJ3p7Mdy5aWCc=")
	loginUrl, err := client.OauthService.GetLoginUrl("http://s.w7.cc")

	if err != nil {
		t.Errorf("message: %s, code: %d", err.Message, err.Errno)
		return
	}
	test := assert.New(t)
	test.NotEmpty(loginUrl)

	fmt.Println(loginUrl)
}

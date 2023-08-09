package service

import (
	"github.com/go-resty/resty/v2"
)

type service struct {
	HttpClient *resty.Client
}

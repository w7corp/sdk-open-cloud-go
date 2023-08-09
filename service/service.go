package service

import (
	"github.com/go-resty/resty/v2"
)

type service struct {
	HttpClient *resty.Client
}

func warpError(message string, errno int) *ErrApiResponse {
	return &ErrApiResponse{
		Message: message,
		Errno:   errno,
	}
}

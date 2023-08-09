package service

type OauthService service

type result struct {
	Url string `json:"url"`
}

func (self *OauthService) GetLoginUrl(redirectUrl string) (string, *ErrApiResponse) {
	apiResult := &result{}
	errResult := &errApiResult{}

	_, err := self.HttpClient.R().
		EnableTrace().
		SetFormData(map[string]string{
			"redirect": redirectUrl,
		}).
		SetResult(apiResult).
		SetError(errResult).
		Post("/we7/open/oauth/login-url/index")

	if err != nil {
		return "", &ErrApiResponse{
			Message: err.Error(),
			Errno:   500,
		}
	}
	if errResult.IsError() {
		return "", errResult.ToError()
	}

	return apiResult.Url, nil
}

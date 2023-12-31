package service

type OauthService service

type ResultAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpireTime  int    `json:"expire_time"`
}

type ResultUserinfo struct {
	UserId         int    `json:"user_id"`
	OpenId         string `json:"open_id"`
	Nickname       string `json:"nickname"`
	Avatar         string `json:"avatar"`
	RoleIdentity   string `json:"role_identity"`
	ComponentAppid string `json:"component_appid"`
	FounderOpenid  string `json:"founder_openid"`
}

func (self *OauthService) GetLoginUrl(redirectUrl string) (string, *ErrApiResult) {
	type result struct {
		Url string `json:"url"`
	}

	apiResult := &result{}
	errResult := &ErrApiResult{}

	_, err := self.HttpClient.R().
		EnableTrace().
		SetFormData(map[string]string{
			"redirect": redirectUrl,
		}).
		SetResult(apiResult).
		SetError(errResult).
		Post("/we7/open/oauth/login-url/index")
	if err != nil {
		return "", NewErrApiResult(err)
	}

	if errResult.IsError() {
		return "", errResult
	}

	return apiResult.Url, NewErrApiResult(nil)
}

func (self *OauthService) GetAccessTokenByCode(code string) (*ResultAccessToken, *ErrApiResult) {
	apiResult := &ResultAccessToken{}
	errResult := &ErrApiResult{}

	_, err := self.HttpClient.R().
		EnableTrace().
		SetFormData(map[string]string{
			"code": code,
		}).
		SetResult(apiResult).
		SetError(errResult).
		Post("/we7/open/oauth/access-token/code")
	if err != nil {
		return nil, NewErrApiResult(err)
	}

	if errResult.IsError() {
		return nil, errResult
	}

	return apiResult, NewErrApiResult(nil)
}

func (self *OauthService) GetUserInfo(accessToken string) (*ResultUserinfo, *ErrApiResult) {
	apiResult := &ResultUserinfo{}
	errResult := &ErrApiResult{}

	_, err := self.HttpClient.R().
		EnableTrace().
		SetFormData(map[string]string{
			"access_token": accessToken,
		}).
		SetResult(apiResult).
		SetError(errResult).
		Post("/we7/open/oauth/user/info")
	if err != nil {
		return nil, NewErrApiResult(err)
	}

	if errResult.IsError() {
		return nil, errResult
	}

	return apiResult, NewErrApiResult(nil)
}

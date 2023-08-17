package service

import "errors"

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

	resp, err := self.HttpClient.R().
		EnableTrace().
		SetFormData(map[string]string{
			"redirect": redirectUrl,
		}).
		SetResult(apiResult).
		SetError(errResult).
		Post("/we7/open/oauth/login-url/index")

	if !resp.IsSuccess() {
		return "", newErrApiResult(errors.New("接口地址错误"))
	}
	if err != nil {
		return "", newErrApiResult(err)
	}
	if errResult.IsError() {
		return "", errResult
	}

	return apiResult.Url, newErrApiResult(nil)
}

func (self *OauthService) GetAccessTokenByCode(code string) (*ResultAccessToken, *ErrApiResult) {
	apiResult := &ResultAccessToken{}
	errResult := &ErrApiResult{}

	resp, err := self.HttpClient.R().
		EnableTrace().
		SetFormData(map[string]string{
			"code": code,
		}).
		SetResult(apiResult).
		SetError(errResult).
		Post("/we7/open/oauth/access-token/code")

	if !resp.IsSuccess() {
		return nil, newErrApiResult(errors.New("接口地址错误"))
	}
	if err != nil {
		return nil, newErrApiResult(err)
	}
	if errResult.IsError() {
		return nil, errResult
	}

	return apiResult, newErrApiResult(nil)
}

func (self *OauthService) GetUserInfo(accessToken string) (*ResultUserinfo, *ErrApiResult) {
	apiResult := &ResultUserinfo{}
	errResult := &ErrApiResult{}

	resp, err := self.HttpClient.R().
		EnableTrace().
		SetFormData(map[string]string{
			"access_token": accessToken,
		}).
		SetResult(apiResult).
		SetError(errResult).
		Post("/we7/open/oauth/user/info")
	if !resp.IsSuccess() {
		return nil, newErrApiResult(errors.New("接口地址错误"))
	}
	if err != nil {
		return nil, newErrApiResult(err)
	}
	if errResult.IsError() {
		return nil, errResult
	}

	return apiResult, newErrApiResult(nil)
}

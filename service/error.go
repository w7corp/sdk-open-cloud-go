package service

import "errors"

func newErrApiResult(err error) *errApiResult {
	if err != nil {
		return &errApiResult{
			ErrMsg: err.Error(),
			Errno:  500,
		}
	} else {
		return &errApiResult{
			ErrMsg: "",
			Errno:  0,
		}
	}
}

type errApiResult struct {
	ErrMsg string `json:"error"`
	Errno  int    `json:"errno"`
}

func (self errApiResult) ToError() error {
	if self.IsError() {
		return errors.New(self.ErrMsg)
	} else {
		return nil
	}
}

func (self errApiResult) IsError() bool {
	return self.ErrMsg != "" || self.Errno > 0
}

func (self errApiResult) Error() string {
	return self.ErrMsg
}

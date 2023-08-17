package service

import "errors"

func newErrApiResult(err error) *ErrApiResult {
	if err != nil {
		return &ErrApiResult{
			ErrMsg: err.Error(),
			Errno:  500,
		}
	} else {
		return &ErrApiResult{
			ErrMsg: "",
			Errno:  0,
		}
	}
}

type ErrApiResult struct {
	ErrMsg string `json:"error"`
	Errno  int    `json:"errno"`
}

func (self ErrApiResult) ToError() error {
	if self.IsError() {
		return errors.New(self.ErrMsg)
	} else {
		return nil
	}
}

func (self ErrApiResult) IsError() bool {
	return self.ErrMsg != "" || self.Errno > 0
}

func (self ErrApiResult) Error() string {
	return self.ErrMsg
}

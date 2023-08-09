package service

type errApiResult struct {
	Error string `json:"error"`
	Errno int    `json:"errno"`
}

func (self errApiResult) ToError() *ErrApiResponse {
	return &ErrApiResponse{
		Message: self.Error,
		Errno:   self.Errno,
	}
}

func (self errApiResult) IsError() bool {
	return self.Error != "" || self.Errno > 0
}

type messageError struct {
	Message string
	Errno   int
	error
}

func (self messageError) Error() string {
	return self.Message
}

type ErrApiResponse messageError
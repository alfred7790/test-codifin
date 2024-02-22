package utils

type DBError struct {
	Code           int
	UserMessage    string
	DevelopMessage error
}

func (e *DBError) Error() string {
	return e.UserMessage
}

func ToUserError(code int, userMsg string, err error) *DBError {
	return &DBError{
		Code:           code,
		UserMessage:    userMsg,
		DevelopMessage: err,
	}
}

func GetCustomError(err error) *DBError {
	if _, ok := err.(*DBError); ok {
		return err.(*DBError)
	} else {
		return &DBError{
			Code:           500,
			UserMessage:    err.Error(),
			DevelopMessage: err,
		}
	}
}

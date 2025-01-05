package errors

const (
	ErrBadRequest = iota + 1
	ErrUserNotFound
	ErrUserExist
	ErrResumeNotFound
	ErrResumeExist
	ErrCustom // This should be the la
)

var definedErrors = map[int]*CustomError{
	ErrBadRequest:     NewCustomError(ErrBadRequest, 400, "Bad Request"),
	ErrUserNotFound:   NewCustomError(ErrUserNotFound, 404, "User not found"),
	ErrUserExist:      NewCustomError(ErrUserExist, 409, "User exists"),
	ErrResumeNotFound: NewCustomError(ErrUserExist, 404, "Resume not found"),
	ErrResumeExist:    NewCustomError(ErrResumeExist, 409, "Resume exists"),
	ErrCustom:         NewCustomError(ErrCustom, 500, "Internal Server Error"),
}

type CustomError struct {
	StatusCode int
	Code       int
	Message    string
}

func NewCustomError(code, status int, message string) *CustomError {
	return &CustomError{
		Code:       code,
		StatusCode: status,
		Message:    message,
	}
}

func (e *CustomError) Error() string {
	return e.Message
}

func GetError(code int) *CustomError {
	return definedErrors[code]
}

func IsCustomError(err error) bool {
	_, ok := err.(*CustomError)
	return ok
}

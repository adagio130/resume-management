package errors

const (
	ErrBadRequest = iota + 1
	ErrUserNotFound
	ErrUserExist

	ErrCustom // This should be the la
)

var definedErrors = map[int]*CustomError{
	ErrBadRequest:   NewCustomError(ErrBadRequest, 400, "Bad Request"),
	ErrUserNotFound: NewCustomError(ErrUserNotFound, 404, "User not found"),
	ErrUserExist:    NewCustomError(ErrUserExist, 409, "User exists"),
	ErrCustom:       NewCustomError(ErrCustom, 500, "Internal Server Error"),
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

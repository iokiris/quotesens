package exceptions

import "errors"

type ErrorType int

type ErrPublic struct {
	Type ErrorType
	Msg  string
	Err  error
}

func (e *ErrPublic) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Msg
} // текст, можно юзеру слать

func (e *ErrPublic) Unwrap() error { return e.Err } // полная ошибка, только внутри

func Err(t ErrorType, msg string, err error) error {
	return &ErrPublic{Type: t, Msg: msg, Err: err}
}

func IsPublic(err error) (ErrorType, string, bool) {
	var pe *ErrPublic
	if errors.As(err, &pe) {
		return pe.Type, pe.Msg, true
	}
	return 500, "", false
}

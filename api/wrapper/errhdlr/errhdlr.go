package errhdlr

import (
	"fmt"
	"reflect"
)

type ErrorStatusMap struct {
	ErrMsg     string
	HTTPStatus int
}

type ErrorMap map[int]ErrorStatusMap

type ErrorWrapper interface {
	GetErrorMap() ErrorMap
}

type AppError struct {
	Code            int
	message         string
	Stack           []string
	isStackRequires bool
	HTTPStatus      int
	RequestID       string
}

func NewAppError(code int, errWrapper ErrorWrapper) *AppError {
	// Caller will make sure the error code
	// will be present in ErrorWrapper
	errMap := errWrapper.GetErrorMap()
	return &AppError{
		Code:            code,
		message:         errMap[code].ErrMsg,
		isStackRequires: false,
		HTTPStatus:      errMap[code].HTTPStatus,
	}
}

func (err *AppError) Error() string {
	standardErr := err.message
	if len(err.Stack) > 0 && err.isStackRequires {
		var temp string
		for i, val := range err.Stack {
			temp += fmt.Sprintf(" %d. %s ", i+1, val)
		}
		return fmt.Sprintf("%s  [%s]", standardErr, temp)
	}
	return standardErr
}

func (err *AppError) EnableTrace() {
	err.isStackRequires = true
}

func (err *AppError) DisableTrace() {
	err.isStackRequires = false
}

func (err *AppError) InjectInternalErrors(errs ...error) {
	for _, val := range errs {
		if val != nil && !reflect.ValueOf(val).IsNil() {
			err.Stack = append(err.Stack, val.Error())
		}
	}
}

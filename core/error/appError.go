package error

import (
    "fmt"
)

type AppError interface {
    GetCode() int
    Error() string
}

type ErrorCodeString struct {
    Msg string
    Code  int
}

func (this *ErrorCodeString) GetCode() int {
    return this.Code
}

func (this *ErrorCodeString) Error() string {
    return fmt.Sprintf("%d - %s", this.Code, this.Msg)
}

func AppErrorNew(code int,msg string) AppError {
    err := &ErrorCodeString{
        msg,
        code,
    }
    return err
}

func Errorf(code int,  format string, arg ... interface{}) AppError {
    msg := fmt.Sprintf(format, arg...)
    return &ErrorCodeString{
        msg,
        code,
    }
}

func ErrorNew(code int,err error) AppError {
    return &ErrorCodeString{
        err.Error(),
        code,
    }
}

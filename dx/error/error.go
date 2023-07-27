package dx_error

import "fmt"

type DxError struct {
	Msg  string
	Code int
}

// 实现 error 接口
func (e *DxError) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

// 实现 error 接口
func Error(msg string, code int) error {
	err := &DxError{
		Msg:  msg,
		Code: code,
	}
	return err
}

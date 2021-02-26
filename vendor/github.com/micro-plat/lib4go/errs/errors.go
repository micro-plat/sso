package errs

import (
	"errors"
	"fmt"

	"github.com/micro-plat/lib4go/types"
)

//ErrNotExist 不存在
var ErrNotExist = errors.New("不存在")

//IError 包含错误码的error
type IError interface {
	Error() string
	GetError() error
	NeedStop() bool
	GetCode() int
	CanIgnore() bool
}

//Error 错误信息
type Error struct {
	code      int
	canIgnore bool
	stop      bool
	error
}

//GetCode 获取错误码
func (a *Error) GetCode() int {
	return a.code
}

//NeedStop 获取结束标记
func (a *Error) NeedStop() bool {
	return a.stop
}

//GetError 获取错误信息
func (a *Error) GetError() error {
	return a
}

//String 格式化错误信息
func (a *Error) String() string {
	return fmt.Sprintf("%d %s", a.code, a.Error())
}

//CanIgnore 是否可以忽略错误
func (a *Error) CanIgnore() bool {
	return a.canIgnore
}

//NewIgnoreError 当前一个可忽略的错误
func NewIgnoreError(code int, err interface{}) *Error {
	ex := NewError(code, err)
	ex.canIgnore = true
	return ex
}

//NewErrorf 创建错误对象
func NewErrorf(code int, f string, args ...interface{}) *Error {
	return NewError(code, fmt.Sprintf(f, args...))
}

//NewError 创建错误对象
func NewError(code int, err interface{}) *Error {
	r := &Error{code: code, canIgnore: false}
	switch v := err.(type) {
	case string:
		r.error = errors.New(v)
	case error:
		r.error = v
	case IError:
		r.error = v.GetError()
	default:
		r.error = errors.New(fmt.Sprint(err))
	}
	return r
}

//NewStop 构建带有结束标记的错误消息
func NewStop(code int, err interface{}) *Error {
	e := NewError(code, err)
	e.stop = true
	return e
}

//NewStopf 构建带有结束标记的错误消息
func NewStopf(code int, f string, args ...interface{}) *Error {
	e := NewErrorf(code, f, args...)
	e.stop = true
	return e
}

//GetCode 获取错误码
func GetCode(err interface{}, def ...int) int {
	switch v := err.(type) {
	case IError:
		return v.GetCode()
	default:
		return types.GetIntByIndex(def, 0, 0)
	}
}

//NeedStop 获取结束标记
func NeedStop(err interface{}, def ...bool) bool {
	switch v := err.(type) {
	case IError:
		return v.NeedStop()
	default:
		return types.GetBoolByIndex(def, 0, false)
	}
}

//GetError 获取错误，当不包含错误时返回空
func GetError(r interface{}) IError {
	switch v := r.(type) {
	case IError:
		return v
	case error:
		return NewError(400, v)
	default:
		return nil
	}
}

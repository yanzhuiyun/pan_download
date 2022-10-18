package logic

import "runtime/debug"

//进行一些错误的包装

type wrapError struct {
	msg   string //出错信息
	stack []byte //栈
	err   error  //原错误
}

func WrapError(err error) *wrapError {
	return &wrapError{
		err:   err,
		stack: debug.Stack(),
	}
}

func WrapErrorWithMsg(err error, msg string) *wrapError {
	return &wrapError{
		msg:   msg,
		err:   err,
		stack: debug.Stack(),
	}
}

func (w *wrapError) Error() string {
	return w.msg + ":  " + w.err.Error()
}

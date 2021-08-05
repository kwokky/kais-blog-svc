package ecode

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"sync"
	"sync/atomic"
)

type Code int

var (
	_messages atomic.Value
	_codes    = map[int]struct{}{}
	once      sync.Once
)

// registerMessage 注册错误信息
func registerMessage() {
	messages := []map[Code]string{PostMessage, CommonMessage}
	newMsg := make(map[Code]string)
	for _, messageMap := range messages {
		for code, msg := range messageMap {
			newMsg[code] = msg
		}
	}

	Register(newMsg)
}

// Register 注册错误码信息
func Register(cm map[Code]string) {
	_messages.Store(cm)
}

// New 生成错误码
func New(e int) Code {
	if e <= 0 {
		panic(fmt.Sprintf("ecode: %d 不能小于等于0", e))
	}

	return add(e)
}

func add(e int) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d 已存在", e))
	}
	_codes[e] = struct{}{}

	return Code(e)
}

type Codes interface {
	// Code get error code
	Code() int
	// Error get error code in string from
	Error() string
	// Message get error message
	Message() string
	// Equal for compatible
	Equal(error) bool
}

func (c Code) Code() int {
	return int(c)
}

func (c Code) Message() string {
	once.Do(func() {
		registerMessage()
	})
	if cm, ok := _messages.Load().(map[Code]string); ok {
		if msg, ok := cm[c]; ok {
			return msg
		}
	}
	return c.Error()
}

func (c Code) Error() string {
	return strconv.FormatInt(int64(c), 10)
}

func (c Code) Equal(err error) bool {
	return EqualError(c, err)
}

func String(e string) Code {
	if e == "" {
		return OK
	}
	// try to int
	code, err := strconv.Atoi(e)
	if err != nil {
		return ServerError
	}
	return Code(code)
}

func Cause(e error) Codes {
	if e == nil {
		return OK
	}

	ce, ok := errors.Cause(e).(Codes)
	if ok {
		return ce
	}

	return String(e.Error())
}

func EqualError(code Codes, err error) bool {
	return Cause(err).Code() == code.Code()
}

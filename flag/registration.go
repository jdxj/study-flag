package flag

import (
	"errors"
)

// CmdHandler 是命令处理器所要实现的接口
type CmdHandler interface {
	Name() string
	Handle() error
}

var ErrRepeatedSubCmd = errors.New("repeated sub commands")
var ErrCmdHandlerNotFound = errors.New("cmd handler not found")

var cmdManager = newSubCmdSet()

func newSubCmdSet() *subCmdSet {
	subCmdSet := &subCmdSet{
		subCmd: make(map[string]CmdHandler),
	}
	return subCmdSet
}

func RegisterCmd(handler CmdHandler) error {
	return cmdManager.registerCmd(handler)
}

func Handle(subCmd string) error {
	handler := cmdManager.handler(subCmd)
	if handler == nil {
		return ErrCmdHandlerNotFound
	}
	return handler.Handle()
}

// subCmdSet 用于方便的注册, 管理命令
type subCmdSet struct {
	subCmd map[string]CmdHandler
}

func (scs *subCmdSet) registerCmd(handler CmdHandler) error {
	if _, ok := scs.subCmd[handler.Name()]; ok {
		return ErrRepeatedSubCmd
	}
	scs.subCmd[handler.Name()] = handler
	return nil
}

func (scs *subCmdSet) handler(subCmd string) CmdHandler {
	return scs.subCmd[subCmd]
}

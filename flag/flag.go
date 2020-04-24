package flag

import (
	"flag"
	"os"
)

// flag.go 文件中的 struct 实现了 CmdHandler 接口, 仅仅是使用例子

type GreetCmd struct {
	// 参数
	say *string
	bye *string

	// 可以组合业务对象
	// abc Abc
}

func (greet *GreetCmd) Name() string {
	return "greet"
}

func (greet *GreetCmd) Handle() error {
	flagSet := flag.NewFlagSet("greet", flag.ExitOnError)

	greet.say = flagSet.String("say", "hello world!", "你想说的话")
	greet.bye = flagSet.String("bye", "Goodbye!", "离开用语!")

	if err := flagSet.Parse(os.Args[2:]); err != nil {
		return err
	}

	// 调用 abc 的方法
	// return abc.Handle(*greet.say, *greet.bye)
	return nil
}

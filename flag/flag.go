package flag

import (
	"flag"
	"fmt"
	"os"
)

// flag.go 文件中的 struct 实现了 CmdHandler 接口, 仅仅是使用例子

type GreetCmd struct {
	say *string
	bye *string
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

	// 一些业务逻辑可能在这里调用
	fmt.Println(*greet.say)
	fmt.Println(*greet.bye)
	return nil
}

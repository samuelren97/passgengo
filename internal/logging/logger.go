package logging

import (
	"fmt"
	"passgengo/internal/colors"
)

func Info(msg string) {
	fmt.Println(colors.Blue+"[INFO]"+colors.Reset,
		"=> "+msg,
	)
}

func Warning(msg string) {
	fmt.Println(
		colors.Yellow+"[WARN]"+colors.Reset,
		"=> "+msg,
	)
}

func Error(err error) {
	fmt.Println(colors.Red+"[ERROR]"+colors.Reset,
		"=> "+err.Error(),
	)
}

func Success(msg string) {
	fmt.Println(colors.Green+"[SUCC]"+colors.Reset,
		"=> "+msg,
	)
}

func Fatal(err error) {
	fmt.Println(colors.Red+"[FATAL]"+colors.Reset,
		"=> "+err.Error(),
	)
	panic(err)
}

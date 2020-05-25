package gen

import (
	"fmt"

	"github.com/ttacon/chalk"
)

func Error(msg string, a ...interface{}) {
	fmt.Println(fmt.Sprint(chalk.Red, fmt.Sprintf(msg, a...), chalk.Reset))
}

func Info(msg string, a ...interface{}) {
	fmt.Println(fmt.Sprint(chalk.Green, fmt.Sprintf(msg, a...), chalk.Reset))
}

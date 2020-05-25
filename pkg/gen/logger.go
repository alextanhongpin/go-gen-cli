package gen

import (
	"fmt"

	"github.com/ttacon/chalk"
)

func Warning(msg string) string {
	return fmt.Sprint(chalk.Red, msg, chalk.Reset)
}

func Success(msg string) string {
	return fmt.Sprint(chalk.Green, msg, chalk.Reset)
}

func Info(msg string) string {
	return fmt.Sprint(chalk.Blue, msg, chalk.Reset)
}

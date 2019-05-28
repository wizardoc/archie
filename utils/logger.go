package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	GREEN = iota + 32
	RED
	YELLOW
	BLUE
)

func sprintColorInfo(conf int, BgColor int, textColor int, text string, directPrint ...bool) string {
	var info = fmt.Sprintf("\n %c[%d;%d;%dm%s%c[0m\n\n", 0x1B, conf, BgColor, textColor, text, 0x1B)

	if len(directPrint) != 0 {
		fmt.Println(info)
	}

	return info
}

func normalColorPrint(info string, textColor int) string {
	return sprintColorInfo(0, 0, textColor, info, true)
}

func Green(info string) string {
	fmt.Println(GREEN)
	return normalColorPrint(info, GREEN)
}

func Red(info string) string {
	return normalColorPrint(info, RED)
}

func Yellow(info string) string {
	return normalColorPrint(info, YELLOW)
}

func Blue(info string) string {
	return normalColorPrint(info, BLUE)
}

func Logger(log string) {
	Green(fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), log))
}

func LogWarn(warn string) {
	fmt.Fprintln(gin.DefaultWriter, Yellow(warn))
}

func LogInfo(info string) {
	fmt.Fprintln(gin.DefaultWriter, Blue(info))
}

func LogError(err error) {
	fmt.Fprintln(gin.DefaultErrorWriter, err)
}

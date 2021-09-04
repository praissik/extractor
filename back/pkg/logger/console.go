package logger

import (
	"fmt"
	"time"
)

func Console(message string) {
	time := time.Now().Format("2006-01-02 15:04:05")
	message = time + " - " + message
	fmt.Println(message)
}

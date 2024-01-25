package errW

import (
	"log"
	"os"
)

func ErrorHandler(str string) {
	flags := log.LstdFlags | log.Lshortfile

	fileErr, _ := os.OpenFile("exchanges/errW/log_err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logErr := log.New(fileErr, "ERROR:\t", flags)
	logErr.Println(str)
}

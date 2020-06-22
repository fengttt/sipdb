package errors

import (
	"fmt"
	"log"
	"os"
)

var logOutput *os.File
func SetLogFile(fn string) {
	var err error
	if logOutput != nil {
		logOutput.Close()
	}

	logOutput, err = os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logOutput)
}

func LogIf(v bool, msg string, args ...interface{}) {
	if v {
		log.Printf(msg, args...)
	}
}

func PanicIf(v bool, k Kind, msg string, args ...interface{}) {
	if v {
		Panic(k, msg, args...)
	}
}

func Panic(k Kind, msg string, args ...interface{}) {
	log.Printf(msg, args...)
	m := fmt.Sprintf(msg, args...)
	e := E(k, m)
	panic(e)
}

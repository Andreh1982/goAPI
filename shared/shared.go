package shared

import (
	"log"
	"os"
	"runtime"
	"strings"
)

var debug bool

func LogCustom(message []string, errorlevel string) {

	if errorlevel == "info" {
		// ZapLogger.Info(message[0])
		Info(message[0])
	} else if errorlevel == "warn" {
		Debug(message[0])
	} else if errorlevel == "error" {
		Error(message[0])
	} else if errorlevel == "fatal" {
		Fatal(message[0])
	}
}

func init() {
	debug = os.Getenv("DEBUG") != ""
}
func Info(msg string, vars ...interface{}) {
	log.Printf(strings.Join([]string{"[INFO]", msg}, " "), vars...)
}
func Debug(msg string, vars ...interface{}) {
	if debug {
		log.Printf(strings.Join([]string{"[DEBUG]", msg}, " "), vars...)
	}
}
func Fatal(msg string) {
	pc, fn, line, _ := runtime.Caller(1)
	if debug {
		log.Fatalf("[FATAL] %s [%s:%s:%d]", msg, runtime.FuncForPC(pc).Name(), fn, line)
	} else {
		log.Fatalf("[FATAL] %s [%s:%d]", msg, fn, line)
	}
}
func Error(msg string) {
	pc, fn, line, _ := runtime.Caller(1)
	if debug {
		log.Printf("[ERROR] %s [%s:%s:%d]", msg, runtime.FuncForPC(pc).Name(), fn, line)
	} else {
		log.Printf("[ERROR] %s [%s:%d]", msg, fn, line)
	}
}

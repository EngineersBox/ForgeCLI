package logging

import (
	"fmt"
	"log"
)

func Info(v ...interface{}) {
	log.Println("[INFO] " + fmt.Sprintln(v...))
}

func Warn(v ...interface{}) {
	log.Println("[WARN] " + fmt.Sprintln(v...))
}

func Error(v ...interface{}) {
	log.Println("[ERROR] " + fmt.Sprintln(v...))
}

func Debug(v ...interface{}) {
	log.Println("[DEBUG] " + fmt.Sprintln(v...))
}

func Fatal(v ...interface{}) {
	log.Println("[FATAL] " + fmt.Sprintln(v...))
}

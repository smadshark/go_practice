package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	fmt.Println(os.Getwd())
	LOGFILE := path.Join("src/log/log_file", "mGo.log")
	logFile, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logFile.Close()

	LstdFlags := log.Ldate | log.Lshortfile
	sLog := log.New(logFile, "LNum ", LstdFlags)
	sLog.Println("Hello there!")

	sLog.SetFlags(log.Lshortfile | log.LstdFlags)
	sLog.Println("Another log entry!")
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var mylogger *log.Logger
	logfile, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	mylogger = log.New(logfile, "\n", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	mylogger.Print("hello")
	mylogger.Print("oh....")
	mylogger.Fatal("test")
	mylogger.Fatal("test2")
}

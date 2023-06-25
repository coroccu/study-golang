package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LogSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func doSomethingOrThrowException() {
	// do something
	panic("Something Wrong!!!!")
}

func doSomething() {
	defer func() {
		s := recover()
		fmt.Println("Recovered: ", s)
	}()
	doSomethingOrThrowException()
}

func main() {
	LogSettings("test.log")
	log.Println("Log")
	log.Printf("%T %v", "Log", "Log")

	// log.Fatalf("%T %v", "Log", "Log")
	// log.Fatalln("Error!") // Not executed

	// file, err := os.Open("something")
	file, err := os.Open("./basic/error_handling/error_handling.go") // You should be in work folder.

	if err != nil {
		log.Fatalln("Error1: ", err)
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error2: ", err)
	}
	fmt.Println(count, string(data))

	//panic & recover
	doSomething()
	fmt.Println("No Panic")
}

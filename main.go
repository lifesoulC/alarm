package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	//	"time"
)

var alarm *Alarm

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: alarm [port]")
		return
	}

	src, e := readIPFile()
	if e != nil {
		fmt.Println(e)
		return
	}
	srcIP = src

	port, e := strconv.Atoi(os.Args[1])
	if e != nil {
		fmt.Println("invalid port number")
		return
	}

	listenPort := fmt.Sprintf(":%d", port)
	fmt.Println("listen port", port)

	var err error
	alarm, err = NewAlarm()
	if err != nil {
		log.Fatal(err)
		return
	}

	e = StartHTTP(listenPort)
	if e != nil {
		fmt.Println(e)
	}
}

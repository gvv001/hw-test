package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func tcpHandle() {

	type clientConfig struct {
		Timeout time.Duration
		Host    string
		Port    string
	}
	var config clientConfig
	flag.StringVar(&config.Host, "host", "localhost", "ip address")
	flag.StringVar(&config.Port, "port", "80", "ip port")
	flag.DurationVar(&config.Timeout, "timeout", 10*time.Second, "request timeout")
	flag.Parse()

	ipAddres := config.Host + ":" + config.Port

	conn, err := net.DialTimeout("tcp", ipAddres, config.Timeout)

	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}

	fmt.Print(conn)

}

func main() {

	fmt.Println("PID:", os.Getpid())

	tcpHandle()
	// c := make(chan os.Signal, 1)

	// signal.Notify(c, syscall.SIGINT)
	// signal.Ignore(syscall.SIGTERM)

	// Place your code here,
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGQUIT)
	//signal.Ignore(syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {

		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true

	}()

	_ = <-done

}

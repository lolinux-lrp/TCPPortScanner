package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func IsValidIP(ipStr string) (valid bool) {
	ip := net.ParseIP(ipStr)
	valid = ip != nil
	return
}

func Connect(host string, port string, timeout time.Duration) error {
	// DialTimeout sets a deadline for the connection attempt
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return err
	}
	if conn != nil { //if conn is nil calling .Close() will give invalid memory address error
		defer conn.Close()
	}

	return nil
}

func ScanIp(ipAddr string) error {

	valid := IsValidIP(ipAddr) //check whether given ip address is valid or not

	if !valid {
		err := fmt.Errorf("Invalid IP address")
		return err
	}

	ch := make(chan struct{}, 100)
	wg := new(sync.WaitGroup)

	for i := 1; i < 65536; i++ {
		wg.Add(1)

		ch <- struct{}{}

		go func(port int) {
			defer wg.Done()
			defer func() { <-ch }()
			err := Connect(ipAddr, strconv.Itoa(port), time.Second)
			if err == nil {
				fmt.Println("Port ", ipAddr, ":", port, " is open.")
			}
		}(i)

	}
	wg.Wait()

	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please enter Ipv4 addr to scan for ports")
	}
	args := os.Args[1]

	err := ScanIp(args)

	if err != nil {
		log.Fatal(err)
	}
}

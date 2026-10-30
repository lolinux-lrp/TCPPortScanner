package main

import (
	"fmt"
	"net"
)

func connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer conn.Close()

	data := []byte("Hello, Server!")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return err
}

func ScanIp(ipAddr string) error {

	for port:=8000; port<8100; port++ {
		portAddr := fmt.Sprintf("%v:%v",ipAddr,port);
		err := connect(portAddr);
		if err!=nil{
			fmt.Println("Error: ",err);
		}else{
			fmt.Println("Port ",portAddr," is open.");
		}
	}
	return nil;
}

func main() {
	err := ScanIp("localhost")
	fmt.Println(err)
}

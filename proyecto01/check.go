package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	addres := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", addres, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v NO es accesible, \n Error: %v ", destination, port)
	} else {
		status = fmt.Sprintf("[UP] %v  ES accesible, \n Peticion desde IP: %v\n Para la IP: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}

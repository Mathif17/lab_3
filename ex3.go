package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	//var serverIP = "10.100.23.147"
	//// var portNumber = ":20021"
	//var serverPort = ":33546"

	/* udp_sender(portNumber)
	udp_reciver(portNumber) */

	//tcpClient(serverIP, serverPort)
	//tcpServer(serverPort)
	go TCP_client("hellop server")
	select {}
}

func udpReciver(rcvPort string) {

	var localIP = "10.100.23.255"

	buffer := make([]byte, 1024)
	ServerAddr, _ := net.ResolveUDPAddr("udp", rcvPort)
	conn, err := net.ListenUDP("udp", ServerAddr)
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		n, addr, _ := conn.ReadFromUDP(buffer)

		if addr.String() != localIP {
			fmt.Println(string(buffer[0:n]))
		}
	}

}

func udpSender(destPort string) {
	var broadcastIP = "255.255.255.255"
	//buffer := make([]byte, 1024)

	local, err := net.ResolveUDPAddr("udp", destPort)
	if err != nil {
		fmt.Println(err)
	}

	remote, err := net.ResolveUDPAddr("udp", broadcastIP+destPort)
	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.DialUDP("udp", local, remote)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello"))
	if err != nil {
		fmt.Println(err)
	}
}

func TCP_client(msg string) {
	addr, _ := net.ResolveTCPAddr("tcp", "10.100.23.147:34933")

	conn, _ := net.DialTCP("tcp", nil, addr)

	go TCP_receiver(conn)

	conn.Write(append([]byte(msg), 0))

	time.Sleep(2 * time.Second)

}

func TCP_receiver(conn *net.TCPConn) {
	buffer := make([]byte, 1024)

	for {
		n, _ := conn.Read(buffer)

		fmt.Println(string(buffer[0:n]))
		time.Sleep(2 * time.Second)
	}
}

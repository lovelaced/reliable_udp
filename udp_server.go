package main

import (
	"fmt"
	"net"
	"os"
//	"bytes"
	"strconv"
)

func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}

func main() {
	serv_addr, err := net.ResolveUDPAddr("udp", ":10001")
	CheckError(err)

	serv_conn, err := net.ListenUDP("udp", serv_addr)
	CheckError(err)

	defer serv_conn.Close()
	read(serv_conn)
}
func read(serv_conn *net.UDPConn) {

	buffer := make([]byte, 1024)

	var last_recd int = 0
	for {
		fmt.Println("Reading from UDP buffer...")
		n,addr,err := serv_conn.ReadFromUDP(buffer)

	//	binary_packet := bytes.NewBuffer(buffer[0:n])

		var current_packet int
		test, err := strconv.Atoi(string(buffer[0:n]))
		current_packet = test

	//	fmt.Println(binary_packet, current_packet)
		go check_packet(serv_conn, current_packet, last_recd)

		fmt.Println("Successfully received", string(buffer[0:n]), "from", addr)
		last_recd, err = strconv.Atoi(string(buffer[0:n]))
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}



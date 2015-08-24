package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
)

func check_packet(serv_conn *net.UDPConn, curr_packet int, last_recd int) {
	fmt.Println("curr_packet:", curr_packet, " last received:", last_recd)
	if curr_packet == last_recd+1 {
		ack(curr_packet)
	} else {
		go ack(curr_packet)
		read(serv_conn) }
}

func ack(curr_packet int) {
	serv_addr,err := net.ResolveUDPAddr("udp","127.0.0.1:10002")
	CheckError(err)

	local_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	conn, err := net.DialUDP("udp", local_addr, serv_addr)
	CheckError(err)

	defer conn.Close()

	msg := strconv.Itoa(curr_packet)

	buffer := []byte(msg)
	fmt.Println("Sending ACK for", curr_packet)
	_, err = conn.Write(buffer)
	if err != nil {
		fmt.Println(msg, err)
	}
	time.Sleep(time.Second * 1)
}

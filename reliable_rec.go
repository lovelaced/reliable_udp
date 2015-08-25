package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
)

func check_packet(serv_conn *net.UDPConn, curr_packet int, last_recd int) {
	// whatever packet we've gotten, we've surely received it, so go ACK it
	go ack(curr_packet)
	// if it's not the next in sequence, we should go back and wait for more
	if curr_packet != last_recd+1 {
		read(serv_conn)
	}
}

func ack(curr_packet int) {
	// initialize outgoing connections for ACKing
	serv_addr,err := net.ResolveUDPAddr("udp","127.0.0.1:10002")
	CheckError(err)

	local_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	conn, err := net.DialUDP("udp", local_addr, serv_addr)
	CheckError(err)

	defer conn.Close()

	// put the packet number we received into the message for sending
	msg := strconv.Itoa(curr_packet)

	buffer := []byte(msg)
	fmt.Println("Sending ACK for", curr_packet)
	// send that packet
	_, err = conn.Write(buffer)
	CheckError(err)

	time.Sleep(time.Second * 1)
}

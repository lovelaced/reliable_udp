package main

import (
	"fmt"
	"net"
	"strconv"
)

// here is where we're constantly monitoring the UDP buffer to wait for those
// precious ACKs we so desperately need to keep the data flowing.
func wait_for_ack(serv_conn *net.UDPConn, expected int, i chan int) {
	for {
		fmt.Println("Waiting for ACK for", expected)
		// our buffer to read into
		buffer := make([]byte, 1024)
		// our wrong value for our ACK result
		ack_val := -1
		n, addr, err := serv_conn.ReadFromUDP(buffer)
		// if we're actually receiving data,
		if addr != nil {
			// do some messy conversions to get the number in the buffer as an int
			curr_packet, err := strconv.Atoi(string(buffer[0:n]))
			CheckError(err)

			fmt.Println("Received", curr_packet, ", looking for", expected)
			// check to see which packet we actually got
			ack_val = check_packet(curr_packet, expected)
			fmt.Println("Received ACK for", curr_packet, "from", addr)
		}
		CheckError(err)
		// if we got an ACK, write it to the channel and get outta there
		if ack_val != -1 {
			i <- ack_val
			break
		}
	}
}

// checks our incoming packet and makes sure we're looking for the right one
func check_packet(curr_packet int, expected int) int {
	if curr_packet == expected {
		return expected
	} else {
		return expected-1
	}
}
package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
)

// starting packet sequence number
var packet_num int = 1

func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
	}
}

func main() {
	// initialize all connections
	serv_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	CheckError(err)
	listen_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10002")
	CheckError(err)
	local_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)
	conn, err := net.DialUDP("udp", local_addr, serv_addr)
	CheckError(err)
	serv_conn, err := net.ListenUDP("udp", listen_addr)
	CheckError(err)
	// create a channel for a packet number to be written to
	i := make(chan int, 1)
	go func () {
		loop:
		// wait for the ack while we're waiting for a packet or timing out
		go wait_for_ack(serv_conn, packet_num, i)
		for {

			select {
			case res := <-i:
				fmt.Println("\nPacket accepted!")
				packet_num = res+1
				// wait for another ack for the next one if we get the right packet
				goto loop
			case <-time.After(100 * time.Millisecond):
				fmt.Println("timed out for", packet_num)
				// if it takes too long for an ACK, go send the packet again
				write(conn)
			}
		}
	}()
	// go write to the connection because the previous stuff is
	// all hanging out in the background for now
	for {
		write(conn)
	}
	defer conn.Close()

}

func write(conn *net.UDPConn) {
	// put the current packet number into the payload
	msg := strconv.Itoa(packet_num)
	// stick it in a buffer
	buffer := []byte(msg)
	fmt.Println("Sending", packet_num)
	// send that packet
	_, err := conn.Write(buffer)
	CheckError(err)
	// chill out for a moment while the magic happens
	time.Sleep(time.Second * 1)

}

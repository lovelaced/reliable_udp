package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)
// keep track of the most recent packet we've accepted
var most_recent int = -1

func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}
//initialize incoming connections
func main() {
	serv_addr, err := net.ResolveUDPAddr("udp", ":10001")
	CheckError(err)

	serv_conn, err := net.ListenUDP("udp", serv_addr)
	CheckError(err)

	defer serv_conn.Close()
	// go read from the connection
	read(serv_conn)
}

func read(serv_conn *net.UDPConn) {

	buffer := make([]byte, 1024)
	// number of last received packet
	var last_recd int = 0
	for {
//		fmt.Println("Reading from UDP buffer...")
		n,addr,err := serv_conn.ReadFromUDP(buffer)

		// get the current packet number
		current_packet, err := strconv.Atoi(string(buffer[0:n]))
		CheckError(err)
		// check the packet we have now against the one we received last
		go check_packet(serv_conn, current_packet, last_recd)
		last_recd, err = strconv.Atoi(string(buffer[0:n]))
		// check for duplicates and discard them as necessary, otherwise mark as received
		if last_recd == most_recent {
			fmt.Println("Duplicate packet found, discarding", string(buffer[0:n]))
		} else {
			fmt.Println("Successfully received", string(buffer[0:n]), "from", addr)
			most_recent = last_recd
		}
		CheckError(err)
	}
}



package main

import (
	"bufio"
	"io"
	"net"
	"time"
)

// OpenUDP to udp network
func OpenUDP(port string) (conn *net.UDPConn, err error) {
	s, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		return
	}

	conn, err = net.ListenUDP("udp4", s)
	if err != nil {
		return
	}

	return
}

// ScanLines from a reader and output to each line
func ScanLines(reader io.Reader, scanned chan<- string, writer io.Writer, remoteOperation <-chan []byte, rwc <-chan io.ReadWriter) (err error) {

	var (
		str       string
		bufReader = bufio.NewReaderSize(reader, 512)
		bufWriter = bufio.NewWriterSize(writer, 1)
		send      []byte
	)

	for {
		select {
		case reader = <-rwc:
			bufReader = bufio.NewReaderSize(reader, 512)
		case send = <-remoteOperation:
			for _, b := range send {
				bufWriter.WriteByte(b)
			}
			bufWriter.WriteByte(13)
		default:
			str, err = bufReader.ReadString('\n')
			if err != nil {
				return
			}
			scanned <- str
		}
		time.Sleep(time.Millisecond)
	}
}

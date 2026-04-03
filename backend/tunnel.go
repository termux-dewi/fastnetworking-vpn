package main

import (
	"io"
	"log"
	"net"
	"github.com/xtaci/kcp-go/v5"
)

func startTunnel() {
	// Mode KCP: NoDelay, Interval 10ms, Resend 2, No Congestion Window
	l, err := kcp.ListenWithOptions(":9000", nil, 10, 3)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handleStream(conn)
	}
}

func handleStream(conn net.Conn) {
	defer conn.Close()
	// Optimasi KCP untuk Latensi Rendah
	sess := conn.(*kcp.UDPSession)
	sess.SetNoDelay(1, 10, 2, 1) 
	sess.SetWindowSize(1024, 1024)
	sess.SetMtu(1350)

	// Proxy logic: Menghubungkan client ke target internet
	// Di sini Anda bisa menambahkan resolver DNS custom
	remote, err := net.Dial("tcp", "8.8.8.8:53") // Contoh DNS forwarding
	if err != nil {
		return
	}
	defer remote.Close()

	go io.Copy(remote, conn)
	io.Copy(conn, remote)
}

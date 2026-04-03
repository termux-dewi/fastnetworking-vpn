package main

import (
	"log"
	"net/http"
)

func main() {
	// Jalankan Data Plane (UDP/KCP) di goroutine
	go startTunnel()

	// Jalankan Control Plane (WebSocket/HTTP)
	http.HandleFunc("/ws", handleMonitor)
	
	log.Println("Superfast Server running...")
	log.Println("Data Port: :9000 (UDP)")
	log.Println("Monitor Port: :8080 (TCP)")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

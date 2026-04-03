package main

import (
 "log"
 "net/http"
)

func main() {
 http.HandleFunc("/tunnel", handleTunnel)
 http.HandleFunc("/ws-monitor", handleMonitor)

 log.Println("🚀 Backend started :8080")
 log.Fatal(http.ListenAndServe(":8080", nil))
}

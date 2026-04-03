package main

import (
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleMonitor(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	for {
		// Simulasi data metrik jaringan
		metrics := map[string]interface{}{
			"ping":   "12ms",
			"jitter": "0.5ms",
			"active_clients": 5,
		}
		ws.WriteJSON(metrics)
		time.Sleep(1 * time.Second)
	}
}

package main

import (
 "net"
 "net/http"
 "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
 CheckOrigin: func(r *http.Request) bool { return true },
}

func handleTunnel(w http.ResponseWriter, r *http.Request) {
 ws, _ := upgrader.Upgrade(w, r, nil)
 defer ws.Close()

 remote, _ := net.Dial("udp", "8.8.8.8:53")
 defer remote.Close()

 done := make(chan bool)

 go func() {
  for {
   _, msg, err := ws.ReadMessage()
   if err != nil { done <- true; return }
   remote.Write(msg)
  }
 }()

 go func() {
  buf := make([]byte, 2048)
  for {
   n, err := remote.Read(buf)
   if err != nil { done <- true; return }
   ws.WriteMessage(websocket.BinaryMessage, buf[:n])
  }
 }()

 <-done
}

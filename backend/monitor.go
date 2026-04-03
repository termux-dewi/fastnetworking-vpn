package main

import (
 "encoding/json"
 "net/http"
 "time"
)

func handleMonitor(w http.ResponseWriter, r *http.Request) {
 ws, _ := upgrader.Upgrade(w, r, nil)
 defer ws.Close()

 start := time.Now()

 for {
  stats := map[string]interface{}{
   "ping": 15,
   "jitter": 1,
   "uptime": time.Since(start).String(),
  }

  msg, _ := json.Marshal(stats)
  ws.WriteMessage(1, msg)
  time.Sleep(time.Second)
 }
}

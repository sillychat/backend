package main

import (
    "fmt"
    "net/http"
    "log"

    "github.com/sillychat/backend/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
  fmt.Println("Websocket Endpoint Hit")
  conn, err := websocket.Upgrade(w, r)
  if err != nil {
    log.Println(w, "%d+V\n", err)
  }
  client := &websocket.Client{
    Conn: conn,
    Pool: pool,
  } 

  pool.Register <- client
  client.Read()
}

func setupRoutes() {
  pool := websocket.NewPool()
  go pool.Start()

  http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    serveWs(pool, w, r)
  })
}

func main() {
    fmt.Println("Chat App v0.01")
    setupRoutes()
    log.Fatal(http.ListenAndServe(":8080", nil))
}

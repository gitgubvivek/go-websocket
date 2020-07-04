package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func httpEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsconnReader(wsconn *websocket.Conn) {
	for {
		msgType, p, err := wsconn.ReadMessage()
		if err != nil {
			log.Error("Failed to read message!")
		}

		log.Info("Received message: ", msgType, p)
	}
}
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Websocket Endppoint")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	wsconn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Unable to upgrade to Websocket Connection!")
	}

	log.Info("Upgraded to Webscket Connection!")

	wsconnReader(wsconn)
}

func setupEndpoints() {
	http.HandleFunc("/", httpEndpoint)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	log.Info("Websocket is awesome!")
	setupEndpoints()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

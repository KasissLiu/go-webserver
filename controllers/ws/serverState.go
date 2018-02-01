package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	m "github.com/KasissLiu/go-webserver/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SyncServerState(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}

	var responMessage []byte

	for {
		messageType := websocket.TextMessage
		state := m.State.GetServerState()
		if responMessage, err = json.Marshal(state); err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, []byte(responMessage)); err != nil {
			log.Println(err)
			return
		}
		time.Sleep(5 * time.Second)
	}

}

package handler

import (
	"fmt"
	"log"
	"net/http"

	"warungpintar/entity"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	AllAPI   = make(map[string]int)
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan entity.Message)
	Total     = 0
)

// Home is a default func in this api webserver
func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("API: Home")
	AllAPI["Home"]++

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome!\n"))
}

// Welcome recieve a string and print it as response
func Welcome(w http.ResponseWriter, r *http.Request) {
	log.Println("API: Welcome")
	name := mux.Vars(r)["name"]
	url := "Welcome?name=" + name
	AllAPI[url]++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Welcome %s!\nThis is a simple api webserver. Don't hope for too much!\n", name)))
}

// History print all api that called before
func History(w http.ResponseWriter, r *http.Request) {
	log.Println("API: History")
	AllAPI["History"]++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for key, val := range AllAPI {
		w.Write([]byte(fmt.Sprintln(key, ":", val)))
	}
}

// Score remember all score that submitted
func Score(w http.ResponseWriter, r *http.Request) {
	log.Println("API: Score")
	AllAPI["Score"]++

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg entity.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("error: ", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

// HandleMessages listen to incoming message and send it to all clients
func HandleMessages() {
	for {
		msg := <-broadcast
		AllAPI["Kill"]++
		Total++
		msg.Kill = fmt.Sprintf("%d", Total)

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("error: ", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

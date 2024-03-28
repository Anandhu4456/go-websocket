package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func SetUpRoutes() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/web", WebHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}
func Reader( conn *websocket.Conn){
	for {
		// p is the message in []byte
		messageType,p,err:=conn.ReadMessage()
		if err!=nil{
			log.Println(err)
			return
		}
		fmt.Println(string(p))
		if err:=conn.WriteMessage(messageType,p);err!=nil{
			log.Println(err)
			return
		}
	}
}
func WebHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type","Application/json")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("client successfully connected...")
	defer conn.Close()
	Reader(conn)
}

func main() {

	SetUpRoutes()
	fmt.Println("server running on port 5000...")
	http.ListenAndServe(":5000", nil)
}

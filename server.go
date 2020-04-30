package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}

func reader(conn *websocket.Conn) {
    for {
        messageType, p , err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }

        log.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }
    }
}

func startSocket(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }
    //fmt.Fprint(w, "Websocket Endpoint")

    websocket, err := upgrader.Upgrade(w, r, nil)
    
    if (err != nil ) {
        log.Println(err)
    }

    log.Println("Websocket Connected")

    reader(websocket)

}

func setupRoutes() {

    /** Websocket endpoints */
    http.HandleFunc("/", homePage)
	http.HandleFunc("/start_socket/", startSocket)

}

func main() {

    fmt.Println("Go websockets")
    setupRoutes()

    // TODO: this needs to be set up as secure (wss) - need to get keys
    // log.Fatal(http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil))
    log.Fatal(http.ListenAndServe(":8081", nil))

}
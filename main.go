package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	http.HandleFunc("/", upgrade)

	go func() { log.Fatal(http.ListenAndServeTLS(":9025", "li.cert", "li.key", nil)) }()
	go func() { log.Fatal(http.ListenAndServeTLS(":9026", "li.cert", "li.key", nil)) }()
	go func() { log.Fatal(http.ListenAndServeTLS(":9027", "li.cert", "li.key", nil)) }()
	go func() { log.Fatal(http.ListenAndServeTLS(":9028", "li.cert", "li.key", nil)) }()
	fmt.Println("Listening on all ports")
	log.Fatal(http.ListenAndServeTLS(":9029", "li.cert", "li.key", nil))
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var dialer = websocket.DefaultDialer

func init() {
	dialer.NetDial = func(network, addr string) (net.Conn, error) {

		return net.Dial("tcp", strings.Replace(addr, "socket.lichess.org", "37.187.205.99", -1))
	}
}

func upgrade(w http.ResponseWriter, r *http.Request) {
	cliconn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	u := r.URL
	u.Scheme = "wss"
	u.Host = r.Host
	r.Header.Del("Sec-Websocket-Key")
	r.Header.Del("Sec-Websocket-Version")
	r.Header.Del("Sec-Websocket-Extensions")
	r.Header.Del("Connection")
	r.Header.Del("Upgrade")

	serconn, _, err := dialer.Dial(u.String(), r.Header)
	if err != nil {
		log.Println(err)
		return
	}

	handleConn(cliconn, serconn)

	//
}

func handleConn(client, server *websocket.Conn) {
	go func() {
		for {
			msg, data, err := server.ReadMessage()
			fmt.Println("S", string(data), err)
			if err != nil {
				client.Close()
				break
			}
			err = client.WriteMessage(msg, data)
			if err != nil {
				client.Close()
				break
			}
		}
	}()
	for {
		msg, data, err := client.ReadMessage()
		fmt.Println("C", string(data), err)
		if err != nil {
			server.Close()
			break
		}
		err = server.WriteMessage(msg, data)
		if err != nil {
			server.Close()
			break
		}
	}

}

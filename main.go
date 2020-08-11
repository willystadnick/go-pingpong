
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

var port string
var player string

func main() {
    http.HandleFunc("/", handler)
    port = os.Args[1]
    player = os.Args[2]
    log.Printf("%s listening on port %s", player, port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    play := r.URL.Path[1:]
    log.Printf("%s responding %s", player, play)
    fmt.Fprintf(w, "%s responding %s", player, play)
    go ping(play)
}

func ping(play string) {
    dest := "player1:8888"
    if (dest == player + ":" + port) {
        dest = "player2:8889"
    }
    resp := "pong"
    if (play == resp) {
        resp = "ping"
    }
    time.Sleep(500 * time.Millisecond)
    _, err := http.Get("http://" + dest + "/" + resp)
    if err != nil {
        log.Fatal(err)
    }
}

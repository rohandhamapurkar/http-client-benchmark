package main

import (
	"io"
	"log"
	"net/http"
)

// SafeCounter is safe to use concurrently.
// type SafeCounter struct {
// 	mu sync.Mutex
// 	v  int
// }

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "This is some response!\n")
		// log.Println("got / request", c.v)
	})

	log.Println("Server Running on localhost:3333")
	err := http.ListenAndServe("127.0.0.1:3333", nil)
	if err != nil {
		log.Fatalln(err)
	}

}

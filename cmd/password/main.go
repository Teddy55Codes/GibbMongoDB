package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
    fileServer := http.FileServer(http.Dir("../../web"))

	http.Handle("/web", fileServer)

    http.HandleFunc("/", indexPage)
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }

}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
			w,
		`<h1>Hello, Gophers</h1>
		<p>You're learning about web development, so</p>
		<p>you might want to learn about the common HTML tags</p>`,
	)
}

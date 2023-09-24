package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("request Host: %s\n", r.RemoteAddr)
	fmt.Printf("request URL: %s\n", r.RequestURI)
	fmt.Printf("request Method: %s\n", r.Method)
	fmt.Print("\nrequest Headers:\n")
	for k, v := range r.Header {
		fmt.Printf("	%s: %s\n", k, v)
	}
	fmt.Println("")
	if r.Method != "GET" {
		fmt.Printf("Request Body: %s\n", r.Body)
	}
	fmt.Println("----------------")
	w.WriteHeader(200)
	io.WriteString(w, "Request Received")

}
func main() {
	var port int
	logPath = flag.String("logpath", "", "Optional filepath to output logs to")
	flag.IntVar(&port, "port", 4434, "Listenting port")
	flag.IntVar(&port, "p", 4434, "Listenting port")

	flag.Parse()

	http.HandleFunc("/", getRequest)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}

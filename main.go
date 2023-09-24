package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Logging struct {
	l *log.Logger
}

func NewLogging(l *log.Logger) *Logging {
	return &Logging{l}
}

func (log *Logging) printRequest(w http.ResponseWriter, r *http.Request) {
	log.l.Printf("Time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	log.l.Printf("request Host: %s\n", r.RemoteAddr)
	log.l.Printf("request URL: %s\n", r.RequestURI)
	log.l.Printf("request Method: %s\n", r.Method)
	log.l.Print("\nrequest Headers:\n")
	for k, v := range r.Header {
		log.l.Printf("	%s: %s\n", k, v)
	}
	log.l.Println("")
	if r.Method != "GET" {
		log.l.Printf("Request Body: %s\n", r.Body)
	}
	log.l.Println("----------------")
	w.WriteHeader(200)
	io.WriteString(w, "Request Received")
}

func main() {
	var port int
	logPath := flag.String("logfile", "", "Optional filepath to output logs to")
	flag.IntVar(&port, "port", 4434, "Listenting port")
	flag.IntVar(&port, "p", 4434, "Listenting port")

	flag.Parse()

	var l *log.Logger

	if *logPath != "" {
		f, err := os.OpenFile(*logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("error openening file :%s\n", err)
			os.Exit(1)
		}
		l = log.New(f, "", 0)
	} else {
		l = log.New(os.Stdout, "", 0)
	}

	lh := NewLogging(l)
	http.HandleFunc("/", lh.printRequest)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}

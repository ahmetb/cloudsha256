package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", sum)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func sum(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "make a POST to this endpoint with data in request body to compute sha256sum")
		return
	}
	log.Printf("received request. content-length: %s", req.Header.Get("content-length"))
	now := time.Now()
	defer func() {
		log.Printf("request completed. took=%s", now)
		req.Body.Close()
	}()

	br := bufio.NewReaderSize(req.Body, 1024*64)
	h := sha256.New()
	io.Copy(h, br)
	fmt.Fprintf(w, "%x", h.Sum(nil))
}

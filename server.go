package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

type httpRequestHandler struct{}

func (h *httpRequestHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	response := []byte("OK")
	responseWriter.Write(response)
}

func main() {
	var httpPort = flag.Int("http-port", 8080, "HTTP Port")
	var profilerHTTPort = flag.Int("profiler-http-port", 6060, "Start profiler localhost")

	// get flags
	flag.Parse()

	// start profiler
	go func() {
		// enable block profiling
		runtime.SetBlockProfileRate(1)

		profilerHTTPAddress := fmt.Sprintf("localhost:%d", (*profilerHTTPort))

		log.Println("Profiler started at address " + profilerHTTPAddress)
		log.Println("Open 'http://" + profilerHTTPAddress + "/debug/pprof/' in you browser or use 'go tool pprof http://" + profilerHTTPAddress + "/debug/pprof/profile' from console")
		log.Println("See details about pprof in https://golang.org/pkg/net/http/pprof/")
		log.Println(http.ListenAndServe(profilerHTTPAddress, nil))
	}()

	httpAddress := fmt.Sprintf("localhost:%d", *httpPort)
	log.Println("HTTP server started at " + httpAddress)

	httpServer := &http.Server{
		Addr:           httpAddress,
		Handler:        &httpRequestHandler{},
		ErrorLog:       log.New(os.Stderr, "", log.LstdFlags),
		ReadTimeout:    time.Duration(5) * time.Second,
		WriteTimeout:   time.Duration(5) * time.Second,
		IdleTimeout:    time.Duration(5) * time.Second,
		MaxHeaderBytes: 1 << 10,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

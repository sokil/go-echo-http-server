package main

import (
	"flag"
)

func main()
{
	var httpPort = flag.Int("http-port", 8080, "HTTP Port")
	var profilerHTTPort = flag.Int("profiler-http-port", 6060, "Start profiler localhost")

	// get flags
	flag.Parse()

	// start profiler
	go func() {
		profilerHTTPAddress := fmt.Sprintf("localhost:%d", *profilerHTTPort)

		log.Println("Profiler started at " + profilerHTTPAddress)
		log.Println("Open 'http://" + profilerHTTPAddress + "/debug/pprof/' in you browser or use 'go tool pprof http://" + profilerHTTPAddress + "/debug/pprof/heap' from console")
		log.Println("See details about pprof in https://golang.org/pkg/net/http/pprof/")
		log.Println(http.ListenAndServe(profilerHTTPAddress, nil))
	}()

	logger := log.New(os.Stderr, "", log.LstdFlags)

	httpAddress := fmt.Sprintf("localhost:%d", httpPort)

	httpServer := &http.Server{
		Addr:           httpAddress,
		Handler:        httpServerHandler,
		ErrorLog:       logger,
		ReadTimeout:    time.Duration(5) * time.Second,
		WriteTimeout:   time.Duration(5) * time.Second,
		IdleTimeout:    time.Duration(5) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
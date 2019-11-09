# run http server with profiling
run-http:
	GOMAXPROCS=1 go run server.go --http-port=8080 -profiler-http-port=6060

# show profiler results in web
run-profiler-web:
	go tool pprof -http=localhost:6061 http://localhost:6060/debug/pprof/profile

run-siege:
	time siege -c 150 -r 250 "http://127.0.0.1:8080"
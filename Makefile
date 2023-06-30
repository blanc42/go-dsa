hello: main.go graph.go deque.go
	go run main.go graph.go deque.go
test: main.go
	go test ./...
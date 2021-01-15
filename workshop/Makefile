build:
	@echo "Building..."
	@go build -o ./server1/main ./server1/main.go
	@go build -o ./server2/main ./server2/main.go
	@go build -o ./server3/main ./server3/main.go
	@go build -o ./server4/main ./server4/main.go

runraw:
	@./server1/main &
	@./server2/main &
	@./server3/main &
	@./server4/main

run:
	make build
	make runraw

run-go: build-go
	./main.exe

build-go:
	go build -o main.exe go/main.go

run-c: build-c
	./a.out

build-c:
	cc c/test.c -lraylib -lGL -lm -lpthread -ldl -lrt -lX11

clean:
	rm a.out

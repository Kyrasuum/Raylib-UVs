run: build
	./a.out

build:
	cc test.c -lraylib -lGL -lm -lpthread -ldl -lrt -lX11

clean:
	rm a.out

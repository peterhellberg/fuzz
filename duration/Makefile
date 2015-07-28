all: run

build:
	go-fuzz-build github.com/peterhellberg/fuzz/duration

clean:
	rm -rf duration-fuzz.zip workdir

run:
	go-fuzz -bin=./duration-fuzz.zip -workdir=workdir

all: run

build:
	go-fuzz-build github.com/peterhellberg/fuzz/duration

clean:
	rm -rf workdir

run:
	go-fuzz -bin=./duration-fuzz.zip -workdir=workdir

all: run

build:
	go-fuzz-build github.com/peterhellberg/fuzz/link

clean:
	rm -rf link-fuzz.zip

run:
	go-fuzz -bin=./link-fuzz.zip -workdir=workdir

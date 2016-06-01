OS  = $(shell uname -s)

ifeq ($(OS), Darwin)
	LDFLAGS += -framework CoreFoundation -framework Security
endif

all:demo

lib:
	- mkdir lib
example_lib:lib
	cd example && go tool cgo example.go && go build -buildmode=c-shared -o ../lib/libexample.so
demo:example_lib
	go build

clean:
	-rm -rf lib golang_dynamic_load_demo example/_obj

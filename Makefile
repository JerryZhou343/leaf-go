.PHONY: all clean wire check

OUTPUT=leaf-go

all: clean wire
	go build  -v -o ./bin/${OUTPUT} cmd/main.go

clean:
	rm -f ./bin/${OUTPUT}

wire:
	wire gen ./di
check:
	wire check ./di
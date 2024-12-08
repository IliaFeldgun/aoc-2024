.PHONY: all output

all: output

output: main input.example
	./$^

main: src/*.go
	go build $^
	mv ./$(shell basename $$PWD) ./main


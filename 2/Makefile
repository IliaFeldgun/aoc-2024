.PHONY: all output

all: output

output: main input
	./$^

main: src/*.go
	go build $^
	mv ./$(shell basename $$PWD) ./main

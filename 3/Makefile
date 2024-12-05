.PHONY: all output

all: output

output: main input
	pwd
	ls
	./$^

main: src/*.go
	go build $^
	mv ./$(shell basename $$PWD) ./main


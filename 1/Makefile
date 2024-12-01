.PHONY: all

all: output

output: 1 input
	./$^

1: src/1.go
	go build $^

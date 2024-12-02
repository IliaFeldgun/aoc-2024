.PHONY: all

all: output

output: 2 input
	./$^

2: src/2.go
	go build $^

.PHONY: all

all: output

output: 3 input
	./$^

3: src/3.go
	go build $^

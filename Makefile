.PHONY: test clean

test: test/1.output test/2.output test/3.output

build/%: %/main
	mkdir -p $(shell dirname $@)
	cp ./$< ./$@
	ls $*

%/main:
	cd $* && $(MAKE)
	ls $*

test/%.output: build/%
	mkdir -p test
	time ./build/$* $*/input > ./$@

clean:
	@rm -rf ./build ./test


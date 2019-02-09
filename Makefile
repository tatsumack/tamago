tamago:
	go build

test: tamago
	./test.sh

clean:
	rm -rf ./tamago tmp*

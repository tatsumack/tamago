tamago:
	go build -o bin/tamago

test: tamago
	./test.sh

clean:
	rm -rf bin/*

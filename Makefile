run:
	go build main.go
	cat ./stdin.txt | ./main

copy:
	rm -rf ./${DIR}
	mkdir ./${DIR}
	cp main.go ./${DIR}
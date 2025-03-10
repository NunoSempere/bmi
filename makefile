build:
	go build -o bmi main.go

run:
	go run main.go

install: bmi 
	sudo mv bmi /usr/bin/bmi

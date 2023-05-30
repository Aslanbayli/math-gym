run:
	go run main.go

build-win:
	go build -o MathGYM.exe main.go 

build-unix:
	go build -o MathGYM.out main.go
	
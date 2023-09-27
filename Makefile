run:
	go run main.go

build:
	rm -rf build
	GOOS=windows GOARCH=amd64 go build -o build/azmena-windows.exe
	GOOS=darwin GOARCH=amd64 go build -o build/azmena-macos
	GOOS=linux GOARCH=amd64 go build -o build/azmena-linux
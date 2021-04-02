build:
		go mod download
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o examples_linux
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o examples_windows
clean:
		go clean
		rm -f examples_linux
		rm -f examples_windows
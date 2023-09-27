build:
	go build -o bin/h-devops main.go

run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"

	# building the program for intel macs
	GOOS=darwin GOARCH=amd64 go build -o bin/h-devops-mac-amd64 main.go
	# building the program for M1 macs
	GOOS=darwin GOARCH=arm64 go build -o bin/h-devops-mac-arm64 main.go
	GOOS=linux GOARCH=arm go build -o bin/h-devops-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/h-devops-linux-arm64 main.go
	# building the program for 64 bits amd/intel linux
	GOOS=linux GOARCH=amd64 go build -o bin/h-devops-linux-amd64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/h-devops-freebsd-386 main.go



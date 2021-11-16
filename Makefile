telecomadmin-superman-linux-amd64:
	@echo "Build telecomadmin-superman-linux-amd64"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "build/telecomadmin-superman-linux-amd64" main.go

telecomadmin-superman-linux-arm:
	@echo "Build telecomadmin-superman-linux-arm"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o "build/telecomadmin-superman-linux-arm" main.go

telecomadmin-superman-linux-arm64:
	@echo "Build telecomadmin-superman-linux-arm64"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o "build/telecomadmin-superman-linux-arm64" main.go

telecomadmin-superman-macos-amd64:
	@echo "Build telecomadmin-superman-macos-amd64"
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o "build/telecomadmin-superman-macos-amd64" main.go

telecomadmin-superman-win-386:
	@echo "Build telecomadmin-superman-win-386"
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o "build/telecomadmin-superman-win-386" main.go

telecomadmin-superman-win-amd64:
	@echo "Build telecomadmin-superman-win-amd64"
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "build/telecomadmin-superman-win-amd64" main.go

telecomadmin-superman-freebsd-amd64:
	@echo "Build telecomadmin-superman-freebsd-amd64"
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o "build/telecomadmin-superman-freebsd-amd64" main.go

all: telecomadmin-superman-linux-amd64 telecomadmin-superman-linux-arm telecomadmin-superman-linux-arm64 telecomadmin-superman-macos-amd64 telecomadmin-superman-win-386 telecomadmin-superman-win-amd64 telecomadmin-superman-freebsd-amd64

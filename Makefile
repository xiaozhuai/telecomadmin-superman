all: telecomadmin-superman-linux-amd64 telecomadmin-superman-linux-arm telecomadmin-superman-linux-arm64 telecomadmin-superman-macos-amd64 telecomadmin-superman-win-386 telecomadmin-superman-win-amd64 telecomadmin-superman-freebsd-amd64


all-zip: telecomadmin-superman-linux-amd64-zip telecomadmin-superman-linux-arm-zip telecomadmin-superman-linux-arm64-zip telecomadmin-superman-macos-amd64-zip telecomadmin-superman-win-386-zip telecomadmin-superman-win-amd64-zip telecomadmin-superman-freebsd-amd64-zip


telecomadmin-superman-linux-amd64:
	@echo "Build exe telecomadmin-superman-linux-amd64"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "build/telecomadmin-superman-linux-amd64" main.go

telecomadmin-superman-linux-amd64-zip: telecomadmin-superman-linux-amd64
	@echo "Build zip telecomadmin-superman-linux-amd64.zip"
	@cd build && zip -FS telecomadmin-superman-linux-amd64.zip telecomadmin-superman-linux-amd64 > /dev/null



telecomadmin-superman-linux-arm:
	@echo "Build exe telecomadmin-superman-linux-arm"
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o "build/telecomadmin-superman-linux-arm" main.go

telecomadmin-superman-linux-arm-zip: telecomadmin-superman-linux-arm
	@echo "Build zip telecomadmin-superman-linux-arm.zip"
	@cd build && zip -FS telecomadmin-superman-linux-arm.zip telecomadmin-superman-linux-arm > /dev/null



telecomadmin-superman-linux-arm64:
	@echo "Build exe telecomadmin-superman-linux-arm64"
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o "build/telecomadmin-superman-linux-arm64" main.go

telecomadmin-superman-linux-arm64-zip: telecomadmin-superman-linux-arm64
	@echo "Build zip telecomadmin-superman-linux-arm64.zip"
	@cd build && zip -FS telecomadmin-superman-linux-arm64.zip telecomadmin-superman-linux-arm64 > /dev/null



telecomadmin-superman-macos-amd64:
	@echo "Build exe telecomadmin-superman-macos-amd64"
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o "build/telecomadmin-superman-macos-amd64" main.go

telecomadmin-superman-macos-amd64-zip: telecomadmin-superman-macos-amd64
	@echo "Build zip telecomadmin-superman-macos-amd64.zip"
	@cd build && zip -FS telecomadmin-superman-macos-amd64.zip telecomadmin-superman-macos-amd64 > /dev/null



telecomadmin-superman-win-386:
	@echo "Build exe telecomadmin-superman-win-386.exe"
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o "build/telecomadmin-superman-win-386.exe" main.go

telecomadmin-superman-win-386-zip: telecomadmin-superman-win-386
	@echo "Build zip telecomadmin-superman-win-386.zip"
	@cd build && zip -FS telecomadmin-superman-win-386.zip telecomadmin-superman-win-386.exe > /dev/null



telecomadmin-superman-win-amd64:
	@echo "Build exe telecomadmin-superman-win-amd64.exe"
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "build/telecomadmin-superman-win-amd64.exe" main.go

telecomadmin-superman-win-amd64-zip: telecomadmin-superman-win-amd64
	@echo "Build zip telecomadmin-superman-win-amd64.zip"
	@cd build && zip -FS telecomadmin-superman-win-amd64.zip telecomadmin-superman-win-amd64.exe > /dev/null



telecomadmin-superman-freebsd-amd64:
	@echo "Build exe telecomadmin-superman-freebsd-amd64"
	@CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o "build/telecomadmin-superman-freebsd-amd64" main.go

telecomadmin-superman-freebsd-amd64-zip: telecomadmin-superman-freebsd-amd64
	@echo "Build zip telecomadmin-superman-freebsd-amd64.zip"
	@cd build && zip -FS telecomadmin-superman-freebsd-amd64.zip telecomadmin-superman-freebsd-amd64 > /dev/null

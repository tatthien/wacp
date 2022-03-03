.PHONY: clean build release

clean: 
	rm -rf ./bin

build: clean
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -o ./bin/wacp.linux-amd64 ./main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -a -o ./bin/wacp.darwin-amd64 ./main.go
	upx --brute bin/wacp.linux-amd64
	upx --brute bin/wacp.darwin-amd64

# Example: make release V=0.0.0
release:
	git tag v$(V)
	@read -p "Please enter to confirm and push to origin ..." && git push origin v$(V)
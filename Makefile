linux:
	@GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 go build -trimpath -tags kqueue -ldflags "-s -w"
MAINFILE = ./cmd/teapot/main.go
OUTFILE = teapot.exe

build:
	go mod tidy
	go build -o ./bin/${OUTFILE} ${MAINFILE}
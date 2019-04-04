cd ../../../
set GOPATH=%cd%
cd src/demo/timewheel_test
set GOARCH=amd64
set GOOS=windows
go build -v -ldflags="-s -w"
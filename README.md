# README

go run ch1/hello
go run ch1/packageNotUsed.go
GODEBUG=gctrace=1 go run gColl.go
time go run sliceGC.go
strace -c go run unsafe.go 2>&1
GOOS=js GOARCH=wasm go build -o main.wasm toWasm.go

~/programs/go/go1.17.2/bin/dlv debug -- 1 2 3
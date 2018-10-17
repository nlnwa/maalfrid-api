package main

// Assumes protoc is installed
//go:generate mkdir -p gen/go
//go:generate protoc -I/usr/include -I. --go_out=plugins=grpc:gen/go maalfrid/service/aggregator/aggregator.proto
//go:generate protoc -I/usr/include -I. --go_out=plugins=grpc:gen/go maalfrid/service/language/ls.proto

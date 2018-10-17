#!/usr/bin/env sh

docker run -v "$(pwd)":/go/src/github.com/nlnwa/maalfrid-api/ grpc/go:1.0 \
/bin/bash -c "go generate github.com/nlnwa/maalfrid-api"

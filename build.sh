#!/usr/bin/env sh

# generate go language proto files using official grpc/go docker image
docker run -v "$(pwd)":/go/src/github.com/nlnwa/maalfrid-api/ grpc/go:1.0 \
/bin/bash -c \
"go generate github.com/nlnwa/maalfrid-api && chmod -R 777 src/github.com/nlnwa/maalfrid-api/gen"

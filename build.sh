#!/usr/bin/env sh

COMMAND=docker
SECURITY_OPT=""
if command -v podman > /dev/null 2>&1; then
    SECURITY_OPT="--security-opt label=disable"
    COMMAND=podman
fi

# generate go language proto files using official grpc/go docker image
${COMMAND} run ${SECURITY_OPT} --rm -v "$(pwd)":/go/src/github.com/nlnwa/maalfrid-api/ grpc/go:1.0 \
/bin/bash -c "go generate github.com/nlnwa/maalfrid-api"

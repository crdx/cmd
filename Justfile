set quiet := true
set shell := ["bash", "-cu", "-o", "pipefail"]

BUILD_DIR := "bin"

[private]
help:
    just --list --unsorted --list-submodules

build name:
    go build -o {{ BUILD_DIR }}/{{ name }} ./{{ name }}/

build-all:
    #!/bin/bash
    set -euo pipefail
    for DIR in */; do
        if [[ -f "${DIR}main.go" ]]; then
            NAME="${DIR%/}"
            echo "Building $NAME"
            go build -o "{{ BUILD_DIR }}/${NAME}" "./${NAME}/"
        fi
    done

clean:
    rm -fv .bin/*

check:
    go build ./...
    go test ./...
    go vet ./...
    golangci-lint run

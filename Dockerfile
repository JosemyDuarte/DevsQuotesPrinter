FROM golang:1.16 AS build-env

ENV APP printer

SHELL ["/bin/bash", "-euf", "-o", "pipefail", "-c"]

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Compilation options:
# - CGO_ENABLED=0: Disable cgo
# - GOOS=linux: explicitly target Linux
# - GOARCH: explicitly target 64bit CPU
# - -trimpath: improve reproducibility by trimming the pwd from the binary
# - -ldflags: extra linker flags
#   - -s: omit the symbol table and debug information making the binary smaller
#   - -w: omit the DWARF symbol table making the binary smaller
#   - -extldflags 'static': extra linker flags: produce a statically linked binary
# - -tags: extra tags
#   - static|static_all|static_build: left out because of no clear documentation on them

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    	-trimpath \
    	-ldflags "-s -w -extldflags '-static'" \
    	-o dist/$APP \
    	./cmd/cli \
 && if ldd dist/$APP; then echo "Should not be dynamically linked binary"; exit 1; fi

FROM alpine:3.9
ENV APP printer

COPY --from=build-env /src/dist/$APP /$APP

#RUN apk add ca-certificates
RUN echo "/$APP \"\$@\"" > /entrypoint.sh && chmod +x /entrypoint.sh

ENTRYPOINT ["/bin/sh","/entrypoint.sh"]
CMD []

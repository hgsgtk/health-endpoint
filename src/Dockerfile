FROM golang:1.12-alpine3.9 as builder

RUN apk add --no-cache bash \
    ca-certificates \
    curl \
    git

ENV GO111MODULE=on

WORKDIR /go/src/github.com/hgsgtk/health-endpoint/src
COPY . /go/src/github.com/hgsgtk/health-endpoint/src

# module cache
# See also https://container-solutions.com/faster-builds-in-docker-with-go-1-11/
COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o server

FROM alpine

COPY --from=builder /go/src/github.com/hgsgtk/health-endpoint/src/server .

ENTRYPOINT ["./server"]
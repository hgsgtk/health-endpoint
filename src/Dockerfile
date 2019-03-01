FROM golang:1.12.0-alpine3.9 as builder

RUN apk add --no-cache bash \
    ca-certificates \
    curl \
    git

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV GO111MODULE=on

WORKDIR /go/src/github.com/hgsgtk/health-endpoint/src
COPY . /go/src/github.com/hgsgtk/health-endpoint/src

RUN go mod download

RUN go build -o server

FROM alpine

COPY --from=builder /go/src/github.com/hgsgtk/health-endpoint/src/server .

CMD ["./server"]
FROM golang:1.20-bookworm

COPY go.mod go.sum /go/src/github.com/sonichigo/letterpress/
WORKDIR /go/src/github.com/sonichigo/letterpress
RUN go mod download
COPY . /go/src/github.com/sonichigo/letterpress
RUN go build -o /usr/bin/letterpress github.com/sonichigo/letterpress/cmd/api

EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/letterpress"]

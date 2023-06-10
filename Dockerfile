FROM golang:latest
RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD . /go/src/work
RUN go install golang.org/x/tools/cmd/godoc@latest
FROM golang:1.14

RUN export GOPATH=/go/src/app
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["prometheus-demo"]
FROM golang:alpine
WORKDIR /go/src/svc_driver
COPY . .
RUN go build -o /go/bin/svc_driver main.go

ENTRYPOINT ["/go/bin/svc_driver"]
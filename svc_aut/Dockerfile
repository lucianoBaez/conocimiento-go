FROM golang:alpine
WORKDIR /go/src/svc_aut
COPY . .
RUN go build -o /go/bin/svc_aut main.go

ENTRYPOINT ["/go/bin/svc_aut"]
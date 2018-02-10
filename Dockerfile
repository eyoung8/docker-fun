FROM golang

RUN mkdir -p /go/src/helloworld
WORKDIR /go/src/helloworld
COPY helloworldserver.go /go/src/helloworld

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -ldflags '-w -s' -a -installsuffix cgo -o helloworld

EXPOSE 8080

CMD ["./helloworld", "8080", "NEW THANG"]

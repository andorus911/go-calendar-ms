FROM golang:latest AS builder
RUN go get -d -v github.com/andorus911/go-calendar-ms/api-ms
WORKDIR /go/src/github.com/andorus911/go-calendar-ms/api-ms
COPY . .
RUN go get -d -v ./... && go install -v ./... && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/andorus911/go-calendar-ms/api-ms .
EXPOSE 58053

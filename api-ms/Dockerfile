FROM golang:latest AS builder
WORKDIR /go/src/go-calendar-ms/api-ms
COPY . .
RUN go get -d -v go-calendar-ms/api-ms && go install -v go-calendar-ms/api-ms && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/go-calendar-ms/api-ms .
EXPOSE 58053

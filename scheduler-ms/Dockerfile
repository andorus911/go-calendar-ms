FROM go-calendar-ms_api AS builder
WORKDIR /go/src/go_learning_homework/go-calendar-ms/scheduler-ms
COPY . .
#RUN go get -d -v go_learning_homework/go-calendar-ms/api-ms/tools/postgres
RUN go get -d -v go_learning_homework/go-calendar-ms/scheduler-ms && go install -v go_learning_homework/go-calendar-ms/scheduler-ms && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/go_learning_homework/go-calendar-ms/scheduler-ms .

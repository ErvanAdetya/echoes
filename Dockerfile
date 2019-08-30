FROM golang:1.11 as builder

EXPOSE 5000

WORKDIR /go/src/github.com/adinb/echoes
RUN go get -d -v github.com/go-chi/chi

COPY main.go .
COPY main_test.go .
RUN go test -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/adinb/echoes/app .

CMD ["./app"]

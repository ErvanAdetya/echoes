FROM golang:1.11

WORKDIR /go/src/github.com/adinb/echoes
RUN go get -d -v github.com/go-chi/chi

COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

CMD ["./app"]

FROM golang:1.12.1-alpine3.9 as builder

WORKDIR /go/src/github.com/nuvo/default-backend
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o main github.com/nuvo/default-backend

FROM scratch

COPY --from=builder /go/src/github.com/nuvo/default-backend/main /
COPY assets/404.html /assets/

CMD ["/main"]
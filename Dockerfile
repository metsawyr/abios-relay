FROM golang:1.19

WORKDIR app

COPY go.mod go.sum ./
RUN go mod download

ADD cmd ./cmd
ADD config ./config
ADD internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -o /abios-relay cmd/app/main.go

EXPOSE 8000
CMD ["/abios-relay"]
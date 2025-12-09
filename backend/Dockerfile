FROM golang:1.25-alpine AS builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o /usr/local/bin/flashcow ./main.go

FROM alpine

COPY --from=builder /usr/local/bin/flashcow /usr/local/bin/flashcow

EXPOSE 80

CMD ["/usr/local/bin/flashcow"]

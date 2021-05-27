FROM golang:alpine as builder
LABEL maintainer="Kamil Juchniewicz <k.juchniewicz@zarna.pl>"
RUN apk update && apk add --no-cache git gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o main . # Wolno kompiluje
RUN CGO_ENABLED=1 go build -a -installsuffix cgo -o main . # Szybko kompiluje

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
EXPOSE 8088
CMD ["./main"]
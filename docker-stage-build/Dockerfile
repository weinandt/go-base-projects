FROM golang:alpine AS builder
WORKDIR /src/
COPY go.mod ./
COPY go.sum ./
COPY server.go ./
RUN go build -o app .

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /src/
COPY --from=builder /src/app ./
RUN ls
CMD ["./app"]  
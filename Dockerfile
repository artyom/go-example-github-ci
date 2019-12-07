FROM golang:alpine AS builder
WORKDIR /app
ENV GOPROXY=https://proxy.golang.org CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go version ; go mod download
COPY . .
RUN go build -ldflags='-s -w' -o main

FROM scratch
COPY --from=builder /app/main /
CMD ["/main"]
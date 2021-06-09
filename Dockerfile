FROM golang:1.16.2 as builder

WORKDIR /build

COPY go.mod go.sum main.go ./
COPY internal/ ./internal
COPY cmd/ ./cmd

ENV GOPRIVATE=pkg.aiocean.dev/*,github.com/aiocean/*
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -v -o cli .

FROM scratch
WORKDIR /root/

COPY --from=builder /build/cli /cli
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/cli"]

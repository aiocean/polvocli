FROM golang:1.16.2 as builder

WORKDIR /build

COPY go.mod go.sum main.go ./
COPY cmd/ ./cmd

ENV GOPRIVATE=pkg.aiocean.dev/*,github.com/aiocean/*
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -v -o polvo .

FROM scratch
WORKDIR /root/

COPY --from=builder /build/polvo /bin/polvo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV PATH "$PATH:/bin"

ENTRYPOINT ["/bin/polvo"]
CMD ["--help"]

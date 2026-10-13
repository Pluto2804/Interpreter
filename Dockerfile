FROM golang:1.25 AS builder

WORKDIR /src

COPY go.mod ./
COPY . .

RUN CGO_ENABLED=0 go build -o /out/steinslang .

FROM alpine:3.22

RUN adduser -D -u 10001 appuser

COPY --from=builder /out/steinslang /usr/local/bin/steinslang

USER appuser

ENTRYPOINT ["steinslang"]

FROM debian:bookworm-slim AS base

FROM golang:1.24.1-bookworm AS builder
COPY --from=base / /
WORKDIR /app
ADD . /app
RUN go build -ldflags '-s -w'

FROM base
COPY --from=builder /app/overscry /
ENTRYPOINT ["./overscry"]

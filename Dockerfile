FROM debian:bookworm-slim AS base

FROM golang:1.25.0-bookworm AS builder
COPY --from=base / /
WORKDIR /app
ADD . /app
RUN make cli-docker

FROM base
COPY --from=builder /app/dist/overscry /
ENTRYPOINT ["./overscry"]

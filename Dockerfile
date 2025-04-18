FROM golang:1.24.2-alpine AS base

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /src
COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download


# Start from the official Golang image
FROM base AS build
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -ldflags="-w -s" -o /out/bin/server ./cmd/*


FROM scratch
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/bin/server /app/server
EXPOSE 8080
CMD ["./server"]

ARG GO_IMAGE=golang:1.26-alpine3.24
ARG RUNTIME_IMAGE=alpine:3.24

FROM ${GO_IMAGE} AS builder

WORKDIR /src

ENV CGO_ENABLED=0 \
    GOCACHE=/go/cache \
    GOFLAGS="-mod=readonly"

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/go/cache \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/go/cache \
    go tool swag init -g cmd/app/main.go -o api/docs/ --parseDependency --parseInternal --useStructName

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/go/cache \
    GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
    go build -trimpath -ldflags="-s -w" -o /app ./cmd/app

FROM ${RUNTIME_IMAGE}

RUN apk add --no-cache ca-certificates \
 && addgroup -S app \
 && adduser -S -G app -u 10001 app

WORKDIR /

COPY --from=builder --chown=app:app /src/config /config

COPY --from=builder --chown=app:app --chmod=0755 /app /usr/local/bin/app

USER app

ENTRYPOINT ["/usr/local/bin/app"]
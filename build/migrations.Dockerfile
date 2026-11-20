FROM alpine:3.24

RUN apk add --no-cache bash ca-certificates wget

ADD https://github.com/pressly/goose/releases/download/v3.27.3/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /migrations

FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN apk add upx git gcc musl-dev tzdata

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN go build -o bin/service cmd/service/main.go
RUN strip /app/bin/service
RUN upx -q -9 /app/bin/service


## Create a minimal docker image
FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/service /bin/service
ENV TZ UTC

CMD [ "/bin/service" ]

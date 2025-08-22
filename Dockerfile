FROM golang:1.24.5-alpine

RUN apk add --no-cache git ca-certificates tzdata curl bash \
    && update-ca-certificates

RUN go install github.com/air-verse/air@latest \
    && cp $(go env GOPATH)/bin/air /usr/local/bin/

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . .
COPY .air.toml .

RUN mkdir -p tmp

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

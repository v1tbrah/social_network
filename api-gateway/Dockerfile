FROM golang:1.20 AS api-gateway

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/

ENV HTTP_HOST=0.0.0.0
ENV HTTP_PORT=7070
ENV LOG_LVL=info

CMD ["/usr/local/bin/app"]

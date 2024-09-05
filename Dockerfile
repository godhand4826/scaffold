FROM golang:1.23.0 AS builder
WORKDIR /builder
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
RUN go build -ldflags '-s -w' -o scaffold .

FROM alpine:3.20.2 AS runner
WORKDIR /app
RUN apk update && apk add --no-cache tzdata
COPY --from=builder /builder/scaffold .
COPY config.yaml .
CMD ["./scaffold"]

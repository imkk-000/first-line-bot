FROM golang:1.12.7 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go get -d github.com/line/line-bot-sdk-go/linebot
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 54321
ENV PORT 54321
ENV CHANNEL_SECRET ${CHANNEL_SECRET}
ENV CHANNEL_TOKEN ${CHANNEL_TOKEN}
CMD ["./app"]

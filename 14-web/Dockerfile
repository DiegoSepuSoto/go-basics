FROM golang:1.13.4 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o main .


FROM alpine:3.10.3  

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"]

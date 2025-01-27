FROM golang:1.21-alpine
WORKDIR /app
RUN go mod init timeservice
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o timeservice
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chown -R appuser:appgroup /app
USER appuser
EXPOSE 8123
CMD ["./timeservice"]

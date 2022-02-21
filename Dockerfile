FROM golang
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build /app/cmd/web && chmod +x /app/cmd/web

ENTRYPOINT ["./web"]

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./web"]
EXPOSE 4000


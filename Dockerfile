FROM golang

COPY . /goapp

WORKDIR /goapp

RUN go mod download

ENV GO111MODULE=on\
    CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64
RUN go build /goapp/web && chmod +x /goapp/web

ENTRYPOINT ["./web"]

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /goapp .

CMD ["./web"]

EXPOSE 4000
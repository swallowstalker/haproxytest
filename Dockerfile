FROM golang:1.14.2-alpine3.11 AS base
WORKDIR /root
COPY . /root/
RUN go build -o simplehttp main.go

FROM alpine:3.11
COPY --from=base /root/simplehttp /bin/simplehttp
ENTRYPOINT ["simplehttp"]
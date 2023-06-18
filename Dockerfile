FROM alpine:3.18.0 AS builder

RUN apk update
RUN apk upgrade
RUN apk add --update go=1.20.4-r0 

WORKDIR /build

ADD . /build

RUN go build main.go

FROM alpine:3.18.0

WORKDIR /opt/caa_tester/

COPY --from=builder /build/main .

CMD ["./main"]
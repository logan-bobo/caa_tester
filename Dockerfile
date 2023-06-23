FROM alpine:3.18.0 AS builder

RUN apk update
RUN apk upgrade
RUN apk add --update \
    bash \
    npm \
    go=1.20.5-r0 

WORKDIR /build

ADD . /build

RUN go build main.go

RUN npm install

RUN ./prepare-static.sh

FROM alpine:3.18.0

WORKDIR /opt/caa_tester/

COPY --from=builder /build/main .
COPY --from=builder /build/static/ ./static
COPY ./templates ./templates

CMD ["./main"]
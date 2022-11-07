FROM golang:1.19-alpine3.16 as builder
ENV CGO_ENABLED=1

RUN apk --no-cache add git ca-certificates wget openssh alpine-sdk build-base gcc zlib-dev

WORKDIR ${GOPATH}/src/github.com/maxwellhealth/action-slack-reporter

COPY . .

RUN go build -o /go/bin/slack-reporter main.go



FROM alpine:3.16

RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/slack-reporter /usr/bin/slack-reporter
CMD [ "slack-reporter" ]

FROM golang:1.19-alpine3.15 as builder

ENV CGO_ENABLED=1
RUN apk --no-cache add git ca-certificates wget openssh alpine-sdk build-base gcc zlib-dev
WORKDIR ${GOPATH}/src/github.com/maxwellhealth/action-slack-reporter
COPY main.go ${GOPATH}/src/github.com/maxwellhealth/action-slack-reporter
ENV GOOS linux
RUN go get -v ./...
RUN go build   -o /go/bin/slack-reporter .



FROM alpine:3.16

RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/slack-reporter /usr/bin/slack-reporter
CMD [ "slack-reporter" ]
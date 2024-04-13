FROM golang:1.22-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata
RUN apk add make

COPY . /banners/

WORKDIR /banners/

RUN make build-server
RUN chown -R root ./bin/banners

FROM alpine:latest


WORKDIR /root/

COPY --from=builder /banners/bin .
COPY --from=builder /banners/configs/banners_config.yml .

CMD [ "./banners", "-config", "banners_config.yml" ]
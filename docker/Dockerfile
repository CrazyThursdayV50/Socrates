FROM golang:1.24.6 AS builder

LABEL stage=gobuilder

ENV GOPROXY=https://goproxy.cn,direct
ENV APP=snake-app

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN go build -ldflags '-linkmode "external" -extldflags "-static"' -o /app/main .

FROM alpine AS final
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/main /app/main
ENV CONFIG_NAME=config
ENTRYPOINT ["/bin/sh", "-c", "./main --d /app/config --n $CONFIG_NAME"]

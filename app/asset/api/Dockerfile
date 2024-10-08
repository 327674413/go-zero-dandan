FROM golang:1.20.5-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY app/asset/api/etc /app/etc
COPY common/land /common/land
RUN go build -ldflags="-s -w" -o /app/asset app/asset/api/asset.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/asset /app/asset
COPY --from=builder /app/etc /app/etc
COPY --from=builder /common/land /common/land
CMD ["./asset", "-f", "etc/asset-api-dev.yaml"]

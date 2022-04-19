FROM golang:1.17.9-alpine3.15 as builder
COPY . /opt/src
WORKDIR /opt/src

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add --update --no-cache gcc musl-dev && \
    CGO_ENABLED=1 GOOS=linux GOPROXY=https://goproxy.cn,direct go build -o /opt/kube-multi-server .

FROM alpine:3.15.4
COPY --from=builder /opt/kube-multi-server /opt/
COPY sql /opt/sql
WORKDIR /opt
ENTRYPOINT ["/opt/kube-multi-server"]

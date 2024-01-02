FROM golang:alpine3.19

#将二进制文件拷贝进容器的GOPATH目录中
COPY music config.toml /go/src/

#为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64 \
GOPROXY="https://goproxy.cn,direct"

#解决docker alpine /bin/sh: ./main: not found
RUN apk add --no-cache libc6-compat

#工作目录
WORKDIR /go/src

#最终运行docker的命令
CMD ["./music"]

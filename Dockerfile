# 使用官方 Golang 镜像作为构建阶段
FROM golang:1.26.3-alpine AS builder
# 设置镜像内的工作目录
WORKDIR /app

# 同时设置大小写代理变量，兼容Go和所有工具
ENV http_proxy=http://192.168.208.213:7897
ENV https_proxy=http://192.168.208.213:7897
ENV HTTP_PROXY=http://192.168.208.213:7897
ENV HTTPS_PROXY=http://192.168.208.213:7897
ENV NO_PROXY=localhost,127.0.0.1,192.168.208.0/24

# 设置Go模块代理和校验服务器
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=gosum.io+ce6e7565+https://goproxy.cn/sumdb

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download -x

# 复制全部项目代码
COPY . .

# 编译二进制文件（输出到 /app/server）
RUN go build -o server ./cmd/server/main.go

# 阶段2：运行环境（超轻量，不带编译工具）
FROM alpine:3.21
WORKDIR /app

# 从编译阶段复制编译好的程序
COPY --from=builder /app/server .
# 复制配置文件
COPY .env .

# 暴露端口
EXPOSE 8080

# 启动命令
CMD ["./server"]

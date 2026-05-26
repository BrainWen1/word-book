# 使用官方 Golang 镜像作为构建阶段
FROM golang:1.26.3-alpine AS builder
# 设置镜像内的工作目录
WORKDIR /app

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

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

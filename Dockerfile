# 基于 Golang 官方镜像构建
FROM golang:1.21.0-alpine3.18 AS builder

# 设置工作目录
WORKDIR /app

# 将本地应用代码复制到容器内的工作目录
COPY . .

# 设置代理、安装依赖、构建二进制文件
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download && go build -o main .

# 运行阶段
FROM alpine:3.18
WORKDIR /app
#复制必要文件到镜像里面
COPY . .
COPY --from=builder /app/main .
CMD ["./main"]
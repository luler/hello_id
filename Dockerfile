# 基于 Golang 官方镜像构建
FROM golang:1.21.0

# 设置工作目录
WORKDIR /app

# 设置代理
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 将本地应用代码复制到容器内的工作目录
COPY . .

# 安装依赖
RUN go mod download

# 构建二进制文件
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# 运行二进制文件
CMD ["./main"]
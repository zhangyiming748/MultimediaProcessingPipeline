FROM golang:1.23.5-alpine3.21 AS builder

# 使用支持 Nvidia GPU 和 CUDA 的基础镜像
FROM nvidia/cuda:11.8.0-devel-ubuntu22.04

# 安装必要的软件包
RUN apt-get update && apt-get install -y \
    ffmpeg \
    git \
    golang-go

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载 Go 模块依赖项
RUN go mod download

# 复制项目源代码
COPY . .

# 编译 Go 程序
RUN go build -o merge merge/merge.go

# 暴露端口（如果需要）
# EXPOSE 8080

# 定义入口点
CMD ["./merge"]
FROM golang:latest

# 设置环境变量
ENV PYTHONWARNINGS="ignore::FutureWarning"

# 标签
LABEL authors="zen"

# 更换国内源
COPY debian.sources /etc/apt/sources.list.d/

# 更新软件并安装依赖
RUN apt update && \
    apt install -y --no-install-recommends \
        python3 \
        python3-pip \
        translate-shell \
        ffmpeg \
        ca-certificates \
        bsdmainutils \
        sqlite3 \
        gawk \
        locales \
        libfribidi-bin \
        dos2unix && \
    rm -rf /var/lib/apt/lists/*

# 安装 openai-whisper 和 yt-dlp
RUN pip install --no-cache-dir openai-whisper yt-dlp

# 复制 Go 程序
WORKDIR /app
COPY . .

# 配置 Go 环境
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct

# 启动程序
ENTRYPOINT ["go", "run", "main.go"]

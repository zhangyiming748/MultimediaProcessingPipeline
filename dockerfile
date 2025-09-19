FROM golang:latest
# docker run -it --rm --name latest golang:latest
# 设置环境变量
ENV PYTHONWARNINGS="ignore::FutureWarning"

# 标签
LABEL authors="zen"

# 更换国内源
COPY debian.sources /etc/apt/sources.list.d/

# 更新软件并安装依赖
RUN apt update
RUN apt install -y --no-install-recommends python3
RUN apt install -y --no-install-recommends python3-pip
RUN apt install -y --no-install-recommends translate-shell
RUN apt install -y --no-install-recommends ffmpeg
RUN apt install -y --no-install-recommends ca-certificates
RUN apt install -y --no-install-recommends bsdmainutils
RUN apt install -y --no-install-recommends sqlite3
RUN apt install -y --no-install-recommends gawk
RUN apt install -y --no-install-recommends locales
RUN apt install -y --no-install-recommends libfribidi-bin
RUN apt install -y --no-install-recommends dos2unix
RUN apt full-upgrade -y
RUN rm -rf /var/lib/apt/lists/*

# 安装 openai-whisper 和 yt-dlp
RUN pip install --no-cache-dir --break-system-packages openai-whisper yt-dlp

# 复制 Go 程序
#WORKDIR /app
#COPY . .

# 配置 Go 环境
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct

# 启动程序
CMD ["bash"]

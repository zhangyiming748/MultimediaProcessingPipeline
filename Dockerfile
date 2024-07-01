FROM golang:1.22.4-bookworm
# docker run -dit -v /Users/zen/Github/WhisperAndTrans:/data --name test golang:1.22.4-bookworm bash
LABEL authors="zen"
# 更换国内源
COPY debian.sources /etc/apt/sources.list.d/
# 更新软件
RUN apt update
RUN apt install -y python3 python3-pip translate-shell ffmpeg ca-certificates bsdmainutils sqlite3 gawk locales libfribidi-bin dos2unix
# 配置pip
# RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
# 安装openai-whisper
RUN rm /usr/lib/python3.11/EXTERNALLY-MANAGED
RUN pip install openai-whisper yt-dlp --no-cache-dir --break-system-packages
# 复制go程序
RUN mkdir /app
WORKDIR /app
COPY . .
# 配置env
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/go/bin
RUN go mod vendor
# 启动程序
ENTRYPOINT ["go", "run","/app/main.go"]
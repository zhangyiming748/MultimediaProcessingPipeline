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
RUN apt install -y --no-install-recommends p7zip-full
RUN apt install -y --no-install-recommends wget
RUN apt install -y --no-install-recommends curl
RUN apt install -y --no-install-recommends build-essential
RUN apt install -y --no-install-recommends mediainfo
RUN apt install -y --no-install-recommends openssh-server
RUN apt install -y --no-install-recommends nano
RUN apt install -y --no-install-recommends axel
RUN apt install -y --no-install-recommends aria2
RUN apt install -y --no-install-recommends htop
RUN apt install -y --no-install-recommends btop
RUN apt install -y --no-install-recommends fonts-wqy-microhei
RUN apt install -y --no-install-recommends fonts-wqy-zenhei
RUN apt install -y --no-install-recommends fonts-noto-cjk
RUN apt full-upgrade -y
RUN rm -rf /var/lib/apt/lists/*

# 安装 openai-whisper 和 yt-dlp
RUN pip install --no-cache-dir --break-system-packages openai-whisper yt-dlp

# 配置 Go 环境
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct

# 设置 root 密码
RUN echo "root:123456" | chpasswd

# 允许 root 登录 SSH
RUN echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config && \
    echo 'PasswordAuthentication yes' >> /etc/ssh/sshd_config

# 重新构建apt缓存
RUN apt update

# 启动 SSH 服务
WORKDIR /
ENTRYPOINT ["service", "ssh", "start", "-D"]

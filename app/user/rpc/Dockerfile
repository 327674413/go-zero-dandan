FROM frolvlad/alpine-glibc:alpine-3.18

# 创建目录
RUN mkdir /user

# 复制文件到容器中
COPY user /user/user
COPY etc/user-dev.yaml /user/user-dev.yaml

# 为可执行文件添加执行权限
RUN chmod +x /user/user

# 设置工作目录
WORKDIR /user

# 设置容器启动时执行的命令
ENTRYPOINT ["./user", "-f", "user-dev.yaml"]

FROM golang:1.18.3

# 项目路径
ENV WORKDIR /root/go-gin

# 传输可执行文件main到镜像
ADD main $WORKDIR/main
ADD config/application.yml $WORKDIR/config/application.yml

# 执行权限命令
RUN chmod +x $WORKDIR/main

# 工作目录
WORKDIR $WORKDIR

# 容器启动时 在工作目录下执行命令
ENTRYPOINT ["./main"]
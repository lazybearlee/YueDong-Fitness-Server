# 声明镜像来源为golang:alpine
FROM golang:alpine
LABEL authors="李卓"

# 声明工作目录
WORKDIR /go/src/fitness

# 拷贝整个server项目到工作目录
COPY . .

# go generate 编译前自动执行代码
# go env 查看go的环境变量
# go build -o server . 打包项目生成文件名为server的二进制文件
RUN go generate && go env && go build -o server .

# 声明镜像来源为alpine:latest
FROM alpine:latest

# 镜像编写者及邮箱
LABEL MAINTAINER="李卓@Lz1958455046@outlook.com"

# 声明工作目录
WORKDIR /go/src/fitness

# 把/go/src/fitness中的可执行文件以及配置文件(resource模板文件)添加进入docker中
COPY --from=0 /go/src/fitness/server ./
COPY --from=0 /go/src/fitness/config.docker.yaml ./
COPY --from=0 /go/src/fitness/resource ./

EXPOSE 8888

# 运行打包好的二进制 并用-c 指定config.docker.yaml配置文件
ENTRYPOINT ./server -c config.docker.yaml


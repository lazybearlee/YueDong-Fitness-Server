FROM centos:7
LABEL authors="李卓"
WORKDIR /opt
ENV LANG=en_US.utf8
COPY deploy/entrypoint.sh .
COPY config.yaml /usr/share/config.yaml
RUN set -ex \
    && echo "LANG=en_US.utf8" > /etc/locale.conf \
    && echo "net.core.somaxconn = 1024" >> /etc/sysctl.conf \
    && echo "vm.overcommit_memory = 1" >> /etc/sysctl.conf \
    && sed -i 's|^mirrorlist=|#mirrorlist=|g' /etc/yum.repos.d/CentOS-Base.repo \
    && sed -i 's|^#baseurl=http://mirror.centos.org|baseurl=http://mirrors.aliyun.com|g' /etc/yum.repos.d/CentOS-Base.repo \
    && yum makecache \
    && yum -y install epel-release \
    && yum -y localinstall http://mirrors.ustc.edu.cn/mysql-repo/mysql57-community-release-el7.rpm \
    && yum -y install mysql-community-server git redis go --nogpgcheck && chmod +x ./entrypoint.sh \
    && go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct \
    && echo "start" > /dev/null

EXPOSE 8888
ENTRYPOINT ["./entrypoint.sh"]

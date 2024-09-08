---
title: Docker Mysql 部署
date: 2024-08-28 16:00:00
categories:
- Docker
- Note
tags:
- Docker
- 部署
- Mysql
- Note
---

## 1. Docker Mysql 部署

本次部署使用的是 Docker Mysql 镜像，该镜像是由官方提供的，可以在 [Docker Hub](https://hub.docker.com/_/mysql) 上查看。

部署采用的版本为 `Mysql`， 即最新版本。

### 1.1 镜像拉取

```bash
docker pull mysql
```

可以采用 `docker images` 查看拉取的镜像。

### 1.2 镜像运行

在镜像运行前，还需要创建一个目录用于挂载到容器中，这样可以保证数据的持久化。

```bash
mkdir -p /mydata/mysql/log
mkdir -p /mydata/mysql/data
mkdir -p /mydata/mysql/conf
```

然后我们创建一个 mysql 容器并指定目录挂载为容器的卷，设置客户端字符集为 `utf8mb4`，并设置 `root` 用户的密码。

```bash
docker run --name mysql -v /mydata/mysql/log:/var/log/mysql -v /mydata/mysql/data:/var/lib/mysql -v /mydata/mysql/conf:/etc/mysql/conf.d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=**** -d mysql:8.0.21 --init-connect="SET collation_connection=utf8mb4_0900_ai_ci" --init-connect="SET NAMES utf8mb4" --skip-character-set-client-handshake
```

- `--name`
   容器名称。
- `-v`
   参数 -v 是 --volume list 的简写，将指定的文件夹挂载为容器的卷（Volume），用来共享文件（日志文件、配置文件、数据文件）。
- `/mydata/mysql/log`
   日志目录。
- `/mydata/mysql/data`
   数据目录。
- `/mydata/mysql/conf`
   配置文件目录。
- `-p 3306:3306`
   参数 -p 是 --publish list 的简写，将3306端口映射到容器的3306端口，对外提供端口。如果同时启动多个mysql容器，对外端口号可以不同，服务之间不会冲突。
- `-e MYSQL_ROOT_PASSWORD=123456`
   参数 -e 是 --env list 的简写，设置环境变量，将 root 用户的密码变量（MYSQL_ROOT_PASSWORD）设置为 123456。
- `-d mysql`
   参数 -d 是 --detach 的简写，指的是容器运行在后台并打印容器ID。后面的mysql可以加版本号，例如mysql:latest、mysql:8.0.31 等等。
- `--init-connect="SET collation_connection=utf8mb4_0900_ai_ci"`
   Client初始化连接Server时，将 collation_connection 排序规则的值设置为 utf8mb4_0900_ai_ci 并作为标志传递给 mysqld 。相当于my.cnf配置文件下，[mysqld]位置下添加 init-connect=“SET collation_connection=utf8mb4_0900_ai_ci” 参数。
- `--init-connect="SET NAMES utf8mb4"`
   Client初始化连接Server时，设置系统变量 NAMES 的值为 utf8mb4 并作为标志传递给 mysqld 。在MySQL官方文档介绍中，设置 NAMES 的字符集，就是给三个会话系统变量 character_set_client、character_set_connection、character_set_results 设置一样的字符集。相当于my.cnf配置文件下，[mysqld]位置下添加 init-connect=“SET NAMES utf8mb4” 参数。
- `--skip-character-set-client-handshake`
   相当于my.cnf配置文件下，【mysqld】位置下添加 skip-character-set-client-handshake 参数。–character-set-client-handshake 的一个开关，用来忽略客户端信息并使用默认服务器字符。官方文档中描述到，在MySQL 4.0版中，服务器和客户端都有一个“全局”字符集，服务器管理员决定使用哪个字符。但MySQL4.1之后，客户端连接服务器时，是有 handshake（握手）的。MySQL4.1之后版本，当客户端连接时，它想发送指定的字符集给服务端来设置 character_set_client、character_set_connection、character_set_results 这三个系统变量，当 mysqld 以 –character-set-server=utf8 这种配置启动时，是无法控制客户端字符集设置的，但MySQL4.0是可以这么做的，为了保留MySQL4.0这种行为，–character-set-client-handshake 开关就诞生了。


### 1.3 进入容器

```bash
docker exec -it mysql bash
```

### 1.4 连接 Mysql

```bash
mysql -uroot -p
```

输入密码即可连接 Mysql 数据库。

### 1.5 修改

```sql
use mysql;
CREATE USER 'lee'@'%' IDENTIFIED BY 'XXXXXX';
GRANT ALL PRIVILEGES ON *.* TO 'lee'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
```

然后还要删除root远程登录的权限，只允许本地登录。

```sql
use mysql;
delete user from mysql.user where user='root' and host='%';
flush privileges;
```

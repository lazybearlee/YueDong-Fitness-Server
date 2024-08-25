#!/bin/bash
if [ ! -d "/var/lib/mysql/fitness" ]; then
    mysqld --initialize-insecure --user=mysql --datadir=/var/lib/mysql
    mysqld --daemonize --user=mysql
    sleep 5s
    mysql -uroot -e "create database fitness default charset 'utf8' collate 'utf8_bin'; grant all on fitness.* to 'root'@'127.0.0.1' identified by '123456'; flush privileges;"
else
    mysqld --daemonize --user=mysql
fi
redis-server &
cd /opt/fitness/ && go run main.go
echo "fitness ALL start!!!"
tail -f /dev/null
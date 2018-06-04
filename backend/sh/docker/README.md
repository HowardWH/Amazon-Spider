# GoSpider-docker

# 一. 介绍

在开发爬虫时，我们要安装好redis，方便做分布式爬虫，装MYSQL，好方便存数据

所以做了这个库，现在是容器化的时代，本人也从事云计算，所以用docker来快速部署环境。

这个库目前部署的环境有`redis`,`mysql`,`golang`。`golang`去除，我建议安装在本地。

你先要有一台Linux操作系统的计算机，我建议安装`ubuntu16.04`。

然后安装好`git`，`docker`和`docker-compose`

请参考：

Dokcer安装：[http://www.lenggirl.com/tool/docker-ubuntu-install.html](http://www.lenggirl.com/tool/docker-ubuntu-install.html)

国内的推荐这样安装：[http://www.lenggirl.com/tool/docker-install.html](http://www.lenggirl.com/tool/docker-install.html)

Docker-compose安装：[http://www.lenggirl.com/tool/docker-compose.html](http://www.lenggirl.com/tool/docker-compose.html)

我已经写好脚本安装docker了，请直接运行

```
chmod 777 docker-install.sh
./docker-install.sh
```

我还是建议你可以百度一下。

# 二. 使用

请先下载该库


```
git clone https://github.com/hunterhug/GoSpider-docker
```


然后赋予`build.sh`执行权限

```
chomd 777 build.sh
./build
```

启动后即可从外部使用。

mysql外部端口8003，账号密码:root/459527502

redis外部端口8002，密码:GoSpider

因为容器挂卷，在内部或外部修改代码，都会同步

如果本机没有安装`mysql`和`redis`客户端，可执行

```
docker exec -it GoSpider-redis redis-cli -a GoSpider
docker exec -it  GoSpider-mysqldb mysql -uroot -p459527502

mysql> show variables like '%max_connect%';
```

进入golang，我建议本地安装，或者你可以通过docker这样安装：

命令如下:

```
docker pull golang:1.8
docker run --rm --net=host -it -v $HOME/mydocker/go:/go --name mygolang golang:1.8 /bin/bash
```

# 三. 原理
`build.sh`内容如下：

```
#!/bin/bash
#sudo rm -rf $HOME/mydocker
sudo mkdir -p $HOME/mydocker/redis/data
sudo mkdir -p $HOME/mydocker/redis/conf
sudo mkdir -p $HOME/mydocker/mysql/data
sudo mkdir -p $HOME/mydocker/mysql/conf
sudo mkdir -p $HOME/mydocker/go
sudo cp my.cnf $HOME/mydocker/mysql/conf/my.cnf
sudo cp redis.conf $HOME/mydocker/redis/conf/redis.conf
sudo docker-compose stop
sudo docker-compose rm -f
sudo docker-compose up -d
```

原理是先将`mysql`和redis`的配置文件移动到根目录下的某个地方，再挂载进容器，数据库数据会保存在本地，即使容器死掉也可重启不丢。

配置文件中`mysql`连接数已经设置高，`redis`设置了密码:`requirepass GoSpider`


`docker-compose.yaml`内容如下：

```
version: '2'
services:
    redis: 
      container_name: "GoSpider-redis"
      image: redis:3.2
      ports: 
        - "6379:6379"
      volumes:
        - $HOME/mydocker/redis/data:/data
        - $HOME/mydocker/redis/conf:/usr/local/etc/redis
      command: redis-server /usr/local/etc/redis/redis.conf
    mysqldb: 
      container_name: "GoSpider-mysqldb"
      image: mysql:5.7
      ports: 
        - "3306:3306"
      environment: 
        - MYSQL_ROOT_PASSWORD=459527502
      volumes:
        - $HOME/mydocker/mysql/data:/var/lib/mysql
        - $HOME/mydocker/mysql/conf:/etc/mysql/conf.d
```

可适当改端口


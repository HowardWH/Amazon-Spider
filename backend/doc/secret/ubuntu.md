# 阿里云ubuntu安装记录

机器配置:

```
root@iZwz9j36enzsqt8ab9tvkxZ:~# lsb_release -a
LSB Version:	core-9.20160110ubuntu0.2-amd64:core-9.20160110ubuntu0.2-noarch:security-9.20160110ubuntu0.2-amd64:security-9.20160110ubuntu0.2-noarch
Distributor ID:	Ubuntu
Description:	Ubuntu 16.04.2 LTS
Release:	16.04
Codename:	xenial

root@iZwz9j36enzsqt8ab9tvkxZ:~# free -h
              total        used        free      shared  buff/cache   available
Mem:           992M        280M         83M        3.1M        628M        534M
Swap:            0B          0B          0B
root@iZwz9j36enzsqt8ab9tvkxZ:~# df -h
Filesystem      Size  Used Avail Use% Mounted on
udev            479M     0  479M   0% /dev
tmpfs           100M  3.1M   97M   4% /run
/dev/vda1        40G  2.7G   35G   8% /
tmpfs           497M  124K  497M   1% /dev/shm
tmpfs           5.0M     0  5.0M   0% /run/lock
tmpfs           497M     0  497M   0% /sys/fs/cgroup
tmpfs           100M     0  100M   0% /run/user/0

```

`Centos`系统操作命令略有不同, 请将`apt`换成`yum`, 其他命令差异请自行解决.

以下是执行过程(全过程, 小白操作):

```
# 登录
ssh root@IP

# 更新安装git
apt update
apt install git

# 安装必要软件
apt install docker.io
apt install docker-compose

# 更新hosts, 复制主机名到hosts
cat /etc/hostname
vim /etc/hosts

>>>
0.0.0.0 iZwz9j36enzsqt8ab9tvkxZ
>>>

# 此后还要保证网络畅通, 提高ulimit的数量, 并且最好添加docker加速器

# 拉代码
mkdir -p ~/gocode/src/github.com/hunterhug
cd ~/gocode/src/github.com/hunterhug
git clone https://github.com/hunterhug/AmazonBigSpider
git clone https://github.com/hunterhug/AmazonBigSpiderWeb

# 启动MYSQL和Redis
cd AmazonBigSpider/sh/docker
chmod 777 ./build.sh
./build

#  检测是否安装成功
docker ps
docker exec -it GoSpider-redis redis-cli -a GoSpider
redis> keys *  (Ctrl+C)

docker exec -it GoSpider-mysqldb mysql -uroot -p459527502
mysql> show databases;
mysql> GRANT ALL PRIVILEGES ON *.* TO 'root'@'%'  IDENTIFIED BY '459527502'  WITH GRANT OPTION;
       flush privileges;
mysql> exit

# scp go1.8压缩包到远程机器
# scp xxxx.tar.gz  ssh@IP:
# 安装golang1.8
# 最新: 直接获取go1.9
wget https://studygolang.com/dl/golang/go1.9.2.linux-amd64.tar.gz
tar -zxvf go1.9.2.linux-amd64.tar.gz
vim /etc/profile.d/myenv.sh

>>>>
export GOROOT=/root/go
export GOPATH=/root/gocode
export GOBIN=$GOPATH/bin
export PATH=.:$PATH:$GOROOT/bin:$GOBIN
:wq
>>>>

source /etc/profile.d/myenv.sh
go env

# 编译爬虫端二进制
cd $GOPATH/src/github.com/hunterhug/AmazonBigSpider
chmod 777 sh/build
sh/build

# 新建数据库
$GOPATH/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/USQL
$GOPATH/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/USQL
$GOPATH/src/github.com/hunterhug/AmazonBigSpider/spiders/de/USQL
$GOPATH/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/USQL

docker exec -it GoSpider-mysqldb mysql -uroot -p459527502
mysql> show databases;
mysql> exit


# 方式一(推荐): 初始化数据库: 使用我抓取好的类目URL, 方式二见最后
cd $GOPATH/src/github.com/hunterhug/AmazonBigSpider/doc/sql
cp * $HOME/mydocker/mysql/conf/

# 类目已经更新了, 现在要导入这个
cp days/usa_category20171026.sql $HOME/mydocker/mysql/conf/

docker exec -it  GoSpider-mysqldb mysql -uroot -p459527502

user jp_smart_base
source /etc/mysql/conf.d/jp_category.sql

user de_smart_base
source /etc/mysql/conf.d/de_category.sql

user smart_base
source /etc/mysql/conf.d/usa_category20171026.sql
#source /etc/mysql/conf.d/usa_category.sql

user uk_smart_base
source /etc/mysql/conf.d/uk_category.sql


# 设置定时器(cdddddddddd@qq.com请换为自己的用户名: 参见:https://proxy.mimvp.com)
crontab -e

>>>>
5 0 * * * ps -ef|grep usa/U* | awk '{print $2}' |xargs -i kill {}
20 0 * * * docker exec -d GoSpider-redis redis-cli -a GoSpider flushall
10 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UURL -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
15 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
20 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
0 3 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
5 0 * * * ps -ef|grep jp/U* | awk '{print $2}' |xargs -i kill {}
20 0 * * * docker exec -d GoSpider-redis redis-cli -a GoSpider flushall
10 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/UURL -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
15 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
20 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
0 3 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
5 0 * * * ps -ef|grep uk/U* | awk '{print $2}' |xargs -i kill {}
20 0 * * * docker exec -d GoSpider-redis redis-cli -a GoSpider flushall
10 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/UURL -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
15 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
20 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
0 3 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
5 0 * * * ps -ef|grep de/U* | awk '{print $2}' |xargs -i kill {}
20 0 * * * docker exec -d GoSpider-redis redis-cli -a GoSpider flushall
10 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/de/UURL -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
15 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/de/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
20 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/de/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
0 3 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/de/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
*/1 * * * * curl http://127.0.0.1:12345/mi?orderid=cdddddddddd@qq.com\&user=jinhan\&password=459527502 > /dev/null 2>&1 &
*/1 * * * * curl http://127.0.0.1:12346/mi?orderid=cdddddddddd@qq.com\&user=jinhan\&password=459527502 > /dev/null 2>&1 &
*/1 * * * * curl http://127.0.0.1:12347/mi?orderid=cdddddddddd@qq.com\&user=jinhan\&password=459527502 > /dev/null 2>&1 &
*/1 * * * * curl http://127.0.0.1:12348/mi?orderid=cdddddddddd@qq.com\&user=jinhan\&password=459527502 > /dev/null 2>&1 &

:wq
>>>>

# 进行测试, 请逐条运行, 真的...

ps -ef|grep usa/U* | awk '{print $2}' |xargs -i kill {}
docker exec -d GoSpider-redis redis-cli -a GoSpider flushall
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UURL -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
ps -ef|grep jp/U* | awk '{print $2}' |xargs -i kill {}
docker exec -d GoSpider-redis redis-cli -a GoSpider flushall
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/UURL -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
ps -ef|grep uk/U* | awk '{print $2}' |xargs -i kill {}
docker exec -d GoSpider-redis redis-cli -a GoSpider flushall
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/UURL -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
ps -ef|grep de/U* | awk '{print $2}' |xargs -i kill {}
docker exec -d GoSpider-redis redis-cli -a GoSpider flushall
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/de/UURL -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/de/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/de/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/de/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

# 看到有12345-12348的端口即可
netstat -ntpl

# 启动网站端
cd $GOPATH/src/github.com/hunterhug/AmazonBigSpiderWeb
go build
./AmazonBigSpiderWeb -s
nohup ./AmazonBigSpiderWeb &
mkdir file
mkdir file/data
mkdir file/back
chmod 777 -R file

打开浏览器输入: /IP:80

# 第二天起就自动了


# 以下需要专业人员进行, 请勿操作(勿动)
# 接着特殊的: 方式二:初始化数据库(包括获取类目URL, 请耐心依次进行, 三个月一次)
# 需要先进数据库删除数据, 请逐行操作, (cdddddddddd@qq.com请换为自己的用户名: 参见:https://proxy.mimvp.com)
cd $GOPATH/src/github.com/hunterhug/AmazonBigSpider/tool/url/
curl "http://127.0.0.1:12345/mi?orderid=cdddddddddd@qq.com&user=jinhan&password=459527502"
curl "http://127.0.0.1:12346/mi?orderid=cdddddddddd@qq.com&user=jinhan&password=459527502"
curl "http://127.0.0.1:12347/mi?orderid=cdddddddddd@qq.com&user=jinhan&password=459527502"
curl "http://127.0.0.1:12348/mi?orderid=cdddddddddd@qq.com&user=jinhan&password=459527502"

docker exec -it GoSpider-mysqldb mysql -uroot -p459527502
>>>
use uk_smart_base
TRUNCATE  table smart_category
use de_smart_base
TRUNCATE  table smart_category
use jp_smart_base
TRUNCATE  table smart_category
use smart_base
TRUNCATE  table smart_category;
>>


# 如果出现很多错误,那你的代理不行, 请将toolproxy设置为false
go run usa_urlmain.go -toolproxy=false -toolstep=0
go run usa_urlmain.go -toolproxy=true -toolstep=1
go run usa_urlmain.go -toolproxy=true -toolstep=2
go run usa_urlmain.go -toolproxy=true -toolstep=3
go run usa_urlmain.go -toolproxy=true -toolstep=4
go run usa_urlparse.go

go run jp_urlmain.go -toolproxy=false -toolstep=0
go run jp_urlmain.go -toolproxy=true -toolstep=1
go run jp_urlmain.go -toolproxy=true -toolstep=2
go run jp_urlmain.go -toolproxy=true -toolstep=3
go run jp_urlmain.go -toolproxy=true -toolstep=4
go run jp_urlparse.go

go run uk_urlmain.go -toolproxy=false -toolstep=0
go run uk_urlmain.go -toolproxy=true -toolstep=1
go run uk_urlmain.go -toolproxy=true -toolstep=2
go run uk_urlmain.go -toolproxy=true -toolstep=3
go run uk_urlmain.go -toolproxy=true -toolstep=4
go run uk_urlparse.go

go run de_urlmain.go -toolproxy=false -toolstep=0
go run de_urlmain.go -toolproxy=true -toolstep=1
go run de_urlmain.go -toolproxy=true -toolstep=2
go run de_urlmain.go -toolproxy=true -toolstep=3
go run de_urlmain.go -toolproxy=true -toolstep=4
go run de_urlparse.go

mv /data/db/usa/url/index.md $GOPATH/src/github.com/hunterhug/AmazonBigSpider
# 导出数据给别人用
cd $GOPATH/src/github.com/hunterhug/AmazonBigSpider/doc/sql/days
docker exec -it GoSpider-mysqldb mysqldump -uroot -p459527502 smart_base smart_category>usa_category$(date +\%Y\%m\%d).sql;
docker exec -it GoSpider-mysqldb mysqldump -uroot -p459527502 de_smart_base smart_category>de_category$(date +\%Y\%m\%d).sql;
docker exec -it GoSpider-mysqldb mysqldump -uroot -p459527502 uk_smart_base smart_category>uk_category$(date +\%Y\%m\%d).sql;
docker exec -it GoSpider-mysqldb mysqldump -uroot -p459527502 jp_smart_base smart_category>jp_category$(date +\%Y\%m\%d).sql;
```


支持`https://www.amazon.com/gp/new-releases`:

```
5 0 * * * ps -ef|grep usa/U* | awk '{print $2}' |xargs -i kill {}
20 0 * * * docker exec -d GoSpider-redis redis-cli -a GoSpider flushall
10 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UURL -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider -toolnew=1 > /dev/null 2>&1 &
15 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider -toolnew=1 > /dev/null 2>&1 &
20 2 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider -toolnew=1 > /dev/null 2>&1 &
0 3 * * * nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider -toolnew=1 > /dev/null 2>&1 &
*/1 * * * * curl http://127.0.0.1:12345/mi?orderid=cdddddddddd@qq.com\&user=jinhan\&password=459527502 > /dev/null 2>&1 &
```


```
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UURL -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider -toolnew=1 > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider -toolnew=1 > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider -toolnew=1 > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider -toolnew=1 > /dev/null 2>&1 &
```
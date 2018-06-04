Centos安装记录:

```
yum update
yum install git
yum install vim
yum install docker

git clone https://github.com/sillybobo/AmazonBigSpider
git clone https://github.com/sillybobo/AmazonBigSpiderWeb

ps -ef|grep usa/U* | awk '{print $2}' |xargs -i kill {}
docker exec -d GoSpider-redis redis-cli -a GoSpider flushall
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UURL -toolnew=1  -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UIP -toolnew=1 -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/ULIST -toolnew=1 -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UASIN -toolnew=1 -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

```

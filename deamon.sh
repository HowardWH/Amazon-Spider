#!/bin/sh
#########美国站#############
ps -fe|grep usa/UIP |grep -v grep
if [ $? -ne 0 ]
then
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

echo `date +"%Y年%m月%d日 %h-%H-%m-%s"` "start usa uip process.....">> /root/restart.log

fi

ps -fe|grep usa/UASIN |grep -v grep
if [ $? -ne 0 ]
then
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

echo `date +"%Y年%m月%d日 %h-%H-%m-%s"` "start usa uasinprocess.....">> /root/restart.log

fi

ps -fe|grep usa/ULIST |grep -v grep
if [ $? -ne 0 ]
then
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

echo `date +"%Y年%m月%d日 %h-%H-%m-%s"` "start usa ulist process.....">> /root/restart.log

fi

########日本######
ps -fe|grep jp/UIP |grep -v grep
if [ $? -ne 0 ]
then
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

echo `date +"%Y年%m月%d日 %h-%H-%m-%s"` "start jp uip process.....">> /root/restart.log

fi

ps -fe|grep jp/UASIN |grep -v grep
if [ $? -ne 0 ]
then
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

echo `date +"%Y年%m月%d日 %h-%H-%m-%s"` "start jp uasinprocess.....">> /root/restart.log

fi

ps -fe|grep jp/ULIST |grep -v grep
if [ $? -ne 0 ]
then
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

echo `date +"%Y年%m月%d日 %h-%H-%m-%s"` "start jp ulist process.....">> /root/restart.log

fi

### 德国 ########
ps -fe|grep de/UIP |grep -v grep
if [ $? -ne 0 ]
then
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/de/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

echo `date +"%Y年%m月%d日 %h-%H-%m-%s"` "start de uip process.....">> /root/restart.log

fi

ps -fe|grep de/UASIN |grep -v grep
if [ $? -ne 0 ]
then
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/de/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

echo `date +"%Y年%m月%d日 %h-%H-%m-%s"` "start de uasinprocess.....">> /root/restart.log

fi

ps -fe|grep de/ULIST |grep -v grep
if [ $? -ne 0 ]
then
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/de/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

echo `date +"%Y年%m月%d日 %h-%H-%m-%s"` "start de ulist process.....">> /root/restart.log

fi

### 英国 ########
ps -fe|grep uk/UIP |grep -v grep
if [ $? -ne 0 ]
then
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/UIP -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

echo `date +"%Y年%m月%d日 %h-%H-%m-%s"` "start uk uip process.....">> /root/restart.log

fi

ps -fe|grep uk/UASIN |grep -v grep
if [ $? -ne 0 ]
then
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/UASIN -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

echo `date +"%Y年%m月%d日 %h-%H-%m-%s"` "start uk uasinprocess.....">> /root/restart.log

fi

ps -fe|grep uk/ULIST |grep -v grep
if [ $? -ne 0 ]
then
nohup /root/gocode/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/ULIST -core=/root/gocode/src/github.com/hunterhug/AmazonBigSpider/public/core -root=/root/gocode/src/github.com/hunterhug/AmazonBigSpider > /dev/null 2>&1 &

echo `date +"%Y年%m月%d日 %h-%H-%m-%s"` "start uk ulist process.....">> /root/restart.log

fi

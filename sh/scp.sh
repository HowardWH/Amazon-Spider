#!/bin/sh

scp -r -v ../spiders/usa/U* root@$1:/root/mydocker/go/src/github.com/hunterhug/AmazonBigSpider/spiders/usa/
scp -r -v ../spiders/jp/U* root@$1:/root/mydocker/go/src/github.com/hunterhug/AmazonBigSpider/spiders/jp/
scp -r -v ../spiders/uk/U* root@$1:/root/mydocker/go/src/github.com/hunterhug/AmazonBigSpider/spiders/uk/
scp -r -v ../spiders/de/U* root@$1:/root/mydocker/go/src/github.com/hunterhug/AmazonBigSpider/spiders/de/

scp -r -v ../config/* root@$1:/root/mydocker/go/src/github.com/hunterhug/AmazonBigSpider/config/

scp -r -v docker root@$1:/root/mydocker/help/
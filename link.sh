#!/bin/bash
mkdir -p $GOPATH/src/github.com/hunterhug
rm $GOPATH/src/github.com/hunterhug/AmazonBigSpider
rm $GOPATH/src/github.com/hunterhug/AmazonBigSpiderWeb

ln -s $(pwd)/backend $GOPATH/src/github.com/hunterhug/AmazonBigSpider
ln -s $(pwd)/webend $GOPATH/src/github.com/hunterhug/AmazonBigSpiderWeb

cd $(pwd)/backend $GOPATH/src/github.com/hunterhug/AmazonBigSpider
godep restore

cd $(pwd)/backend $GOPATH/src/github.com/hunterhug/AmazonBigSpiderWeb
godep restore

chmod 777 $GOPATH/src/github.com/hunterhug/AmazonBigSpider/sh/docker/build.sh
$GOPATH/src/github.com/hunterhug/AmazonBigSpider/sh/docker/build.sh

#chmod 777 $GOPATH/src/github.com/hunterhug/AmazonBigSpider/sh/build.sh
#$GOPATH/src/github.com/hunterhug/AmazonBigSpider/sh/build.sh

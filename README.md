# Full Automatic Amazon Distributed crawler|spider (USA, Japan, Germany and UK)

[![GitHub forks](https://img.shields.io/github/forks/hunterhug/AmazonBigSpider.svg?style=social&label=Forks)](https://github.com/hunterhug/AmazonBigSpider/network)
[![GitHub stars](https://img.shields.io/github/stars/hunterhug/AmazonBigSpider.svg?style=social&label=Stars)](https://github.com/hunterhug/AmazonBigSpider/stargazers)
[![GitHub last commit](https://img.shields.io/github/last-commit/hunterhug/AmazonBigSpider.svg)](https://github.com/hunterhug/AmazonBigSpider)
[![Go Report Card](https://goreportcard.com/badge/github.com/hunterhug/AmazonBigSpider)](https://goreportcard.com/report/github.com/hunterhug/AmazonBigSpider)
[![GitHub issues](https://img.shields.io/github/issues/hunterhug/AmazonBigSpider.svg)](https://github.com/hunterhug/AmazonBigSpider/issues)

此项目为亚马逊电商带来了过多压力， 故开始闭源， 新的功能和修的新BUG将不再提交Github. 本人不承担该数据采集程序所带来的一切纠纷， 禁止任何机构及个人将此系统作为商业用途！ 

## 一. 介绍

### 1. 中文介绍

用途： 选款，特别适合亚马逊跨境电子商务运营公司(不支持中国亚马逊)

核心竞争力： 四个国际站点：`美国/英国/日本/德国`，分布式，配套后台可视化

原理： 首先通过获取亚马逊所有类目的URL，即从第一层大类，一直获取到第六层小类。通过这些类目URL可以依次抓取到这些类目某段时间的Top100的商品（类目下的爆款），这些Top100的商品排名我们称为小类排名，每个小时会变一次，但是由于变化基本不会太频繁以及抓取的商品数量很多，基本能覆盖。比如：有一个大类，下面有某一个三层类目，这个三层类目下面有几十个四层，四层下面又有五层，很多个Top100组在一起构成了三层我们需要的商品。通过这些小类商品数据，我们再进详情页获取更多的字段（评论数，星数，是否FBA，价格等），包括每件商品的最顶层排名，我们称大类排名。通过商品去重，分布式代理以及数据的一些预处理设计，加大马力，运用IT采集技术，我们能得到亚马逊大部分卖得好的商品，通过筛选，排序，我们可以从不同角度观察商品趋势。对于卖家来选款的话是极好的。

关于选款： 亚马逊和国内天猫的差别在于店铺概念弱化，亚马逊以单品为为单位，基本一个ASIN就是一个商品类型，卖得好的商品很多人可以跟卖。不同的商家会有一样ASIN的商品，如果谁的商品好（省略...）。步骤一般是：通过该平台Web端查看某大类排名前一万名，进行一些筛选，比如价格在20刀的，FBA的商品，然后可以再点进去商品，看这件商品十几天的排名和价格变化等，然后我决定跟卖，先去阿里巴巴批发看看有没有这个东西，有！价格利润很多。好，我们卖！然后每天可以上来平台搜我们这件商品的ASIN，查看最近的变化。

因为这套产品包括爬虫端和网站端（可视化数据，筛选导出数据，结果见result结果文件夹），并且对技能要求较高，安装并稳定运行较复杂，可见[产品概况](https://github.com/hunterhug/marmot/blob/master/doc/amazon.md)，彻底开源. 这是BI爬虫端, BI网站端见： [亚马逊四站BI产品网站端](https://github.com/hunterhug/AmazonBigSpiderWeb)

英文已经凌乱，Old English Read this [OutOfDateYouShouldReferChinese](old-readme.md) 仔细阅读，有益身心.

亚马逊爬虫支持

1. 列表页和详情页可选择代理方式
2. 多浏览器保存cookie机制
3. 机器人检测达到阈值自动换代理
4. 检测日期过期自动停止程序
5. IP池扫描周期填充代理IP
6. 支持分布式跨平台抓取
7. 高并发进程设置抓取
8. 默认网页爬取去重
9. 日志记录功能
10. 配套可视化网站，支持多角度查看数据，小类数据，大类数据，Asin数据和类目数据，支持查看每件Asin商品的历史记录，如排名，价格，打分，reviews变化。部分数据支持导出，且网站支持RBAC权限，可分配每部分数据的查看和使用权限。
11. 网络端监控爬虫，可查看爬虫当前时段数据抓取状态，爬取的进度，IP的消耗程度。   **可支持网络端启动和停止爬虫，彻底成为Saas**（待做）
12. 可自定义填入IP，如塞入其他代理IP网站API获取的IP
13. 可选择HTML文件保存本地

分布式，高并发，跨平台，多站点，多种自定义配置，极强的容错能力是这个爬虫的特点。机器数量和IP代理足够情况下，每天每个站点可满足抓取几百万的商品数据。

### 2. 英语介绍

Support UAS/Japan/Germany/UK, Amazing!

Catch the best seller items in Amazon USA! Using redis to store proxy ip and the category url. First fetch items list and then collect many Asin, store in mysql. Items list catch just for the Asin, and we suggest one month or several weeks to fetch list page. We just need fetch the Asin detail page and everything we get!

We keep all Asin in one big table. And if catch detail 404, we set it as not valid. Also we can use API catch big rank but look not so good!

So, there are two ways to get the big rank：

1. catch list page(not proxy), using API get the big rank

2. catch list page(not proxy), and then get asin detail page(proxy), API can not catch all the asin big rank so must use this!

Due to we want list smallrank and the bigrank at the same time, but mysql update is so slow, we make two tables to save, one is smallrank, one is bigrank!

We test a lot,if a ip no stop and more than 500 times http get,the list page will no robot,but the detail asin page will be robot. So we bind a proxy ip with and fix useragent, and keep all cookie. But it still happen, a IP die still can fetch detail page after 26-100times get, It tell us we can still ignore robot, and catch max 100 times we will get that page. robot page is about 7KB.

However, if a lot of request, will be like that 500 error

For reason that the detail page is such large that waste a lot of disk space, we save the list page in the local file and the detail page you can decide whether to save it or not.


## 3. 最新说明

鉴于本人精力有限, 无暇多开发新功能, 有更多需求可来邮. 目前搭了一套[亚马逊电子商务大数据智能平台](http://aws.lenggirl.com), 您可以上去观摩, 帐号密码均为`admin`, 切勿破坏, 且行且珍惜. 如果您是一名开发, 您觉得不错, 学习到了知识, 可以扫描下方二维码友情赞助. 如果你是一个电商服务公司的老板, 或者是从业者, 急需使用到该平台来进行选款, 洞察商品变化趋势, 可以来邮咨询, 我提供有偿搭建服务, 价格合理, 完全划得来.

核心的爬虫包也已经拆分成库了，见[Project:Marmot(Tubo) - Golang Web Spider/Crawler/Scrapy Package | 爬虫库](https://github.com/hunterhug/marmot)。网站端也拆分成库了[Project:Rabbit(Tuzi) - Golang Enterprise Web | 简单企业网站](https://github.com/hunterhug/rabbit)

如果这个产品有帮助到你,可以抛出请我吃下辣条吗?

微信
![微信](https://raw.githubusercontent.com/hunterhug/hunterhug.github.io/master/static/jpg/wei.png)

支付宝
![支付宝](https://raw.githubusercontent.com/hunterhug/hunterhug.github.io/master/static/jpg/ali.png)

爬虫有风险, 本人不承担由此开源项目带来的任何责任。

## 4. 版本说明

v2.0

1. 增加安装详细说明
2. 修补一些BUG
3. 美国站类目URL已经更新: /doc/sql/days/usa_category20171026.sql(数据库导入必须是最新的)

v2.3

1. 解决许多BUG

以下为安装使用文档

## 二. 文件目录

```
├── config  配置文件：运行前必须配置
│   ├── de_config.json    德国亚马逊爬虫远程配置（在本地新建一个空文件：`远程.txt` 即加载此配置）
│   ├── de_local_config.json   德国亚马逊爬虫本地配置（默认加载这个, 以下是不同站点的配置）
│   ├── de_log.json	德国亚马逊日志记录文件
│   ├── jp_config.json 
│   ├── jp_local_config.json
│   ├── jp_log.json
│   ├── uk_config.json
│   ├── uk_local_config.json
│   ├── uk_log.json
│   ├── usa_config.json
│   ├── usa_local_config.json
│   └── usa_log.json
├── doc
│   ├── categoryrecord.xlsx   你可以看看四站（ 美国/日本/英国/德国的商品类目情况）
│   ├── img
│   └── sql   你必须手动导入的四站类目SQL，很大，见tool/url
├── public
│   ├── core   核心包
│   └── log
├── result结果
│   ├── 20170731Clothing.xlsx   从网站端导出的数据(示例数据)
│   └── 20170731Kitchen&Dining.xlsx
├── sh          这个是脚本，我们用来快速启动爬虫的
│   ├── docker   用来快速启动docker版redis和mysql
│   ├── build.sh 在本地编译二进制，然后直接通过以下scp.sh发送到阿里云等机器
│   ├── scp.sh
│   ├── de-crontab.txt  定时器
│   ├── jp-crontab.txt
│   ├── uk-crontab.txt
│   ├── usa-crontab.txt
├── spiders   这个是爬虫入口
│   ├── de
│   ├── jp
│   ├── uk
│   └── usa
└── tool
    ├── python
    └── url   四站类目数据爬取程序在这里，需要手工改代码做类目，每隔个两三个月就需要重爬一次(确认大类十分复杂...)
├── ip.txt   你可以将固定的代理IP放在这里，因为亚马逊详情页爬太多会反爬虫

```

## 三. 如何使用

安装十分复杂, 可来邮咨询.

### 1. 获取代码/安装环境

省略...

### 2. 配置爬虫

省略...

### 3. 编译程序

爬虫入口在：

```
├──spiders
	├── de
	│   ├── asinmain.go
	│   ├── asinpool.go
	│   ├── initsql.go
	│   ├── ippool.go
	│   ├── listmain.go
	│   ├── listparsemain.go
	│   └── urlpool.go
	├── jp
	│   ├── asinmain.go
	│   ├── asinpool.go
	│   ├── initsql.go
	│   ├── ippool.go
	│   ├── listmain.go
	│   ├── listparsemain.go
	│   └── urlpool.go
	├── uk
	│   ├── asinmain.go
	│   ├── asinpool.go
	│   ├── initsql.go
	│   ├── ippool.go
	│   ├── listmain.go
	│   ├── listparsemain.go
	│   └── urlpool.go
	└── usa
	    ├── asinmain.go  4. 抓取详情页，补充大类排名等商品信息，打Mysql大类数据和Hash方便查看历史趋势
	    ├── asinpool.go  中间产物,不用
	    ├── initsql.go  1.初始化数据库
	    ├── ippool.go   2.插代理IP到Redis并监控爬虫
	    ├── listmain.go	4.抓取类目列表Top100，打redis记录额外数据以及打Mysql小类数据
	    ├── listparsemain.go 中间产物,不用
	    └── urlpool.go  3.打类目URL到redis，供4步骤使用
```

省略...

### 4. 初始化数据库

如果不申明，都是以美国站为例。需要填充四个站点8个数据基本数据库，以及4*80=320个HASH库，要运行上面编译好的二进制, 执行:


省略...

```
doc
└── sql
    ├── de_category.csv  德国的类目CSV，你可以打开看看
    ├── de_category.sql  需要导入的数据库
    ├── jp_category.csv
    ├── jp_category.sql
    ├── uk_category.csv
    ├── uk_category.sql
    ├── usa_category.csv
    └── usa_category.sql
```

省略...


### 5. 运行程序

省略...

### 6. 如何使用代理IP

因为亚马逊四站点对详情页会反爬虫，一个IP可以抓500页然后被封，但是现在市场上卖的代理IP特别多，几百万动态那种，所以不用担心这个问题。如果你自己有固定的代理IP，请把它写在`ip.txt`里面。

自建代理请见[多IP多网关Squid架设Http服务器](http://www.lenggirl.com/tool/overwall.html#http)

省略...

### 7. 分布式部署(可选)

分布式部署时，由于数据量巨大，开启网站端时，容易卡，所以你可以对数据库进行读写分离，一般数据量 `不大` 可以不用。

可以购买腾讯云，阿里云，亚马逊云，我们一般只需买一台，花费大概是每年1500左右，然后根据上述起`docker mysql/redis`（你也可以自行安装），然后在本地编译好程序，使用`./scp.sh 189.55.55.55`将二进制文件以及配置文件，传到远程机器，跑一下测试，测试没问题，再开启定时器。`scp.sh`在`sh`文件夹中。

省略...


### 8. 网站端

BI产品爬虫端的价值大，但是配套网站端，价值可以翻好多倍。你可以查看[Full Golang Automatic Amazon Distributed crawler|spider (USA, Japan, Germany and UK) | 亚马逊四站BI产品网站端 ](https://github.com/hunterhug/AmazonBigSpiderWeb)进行安装。

截图如下：

类目，你可以自行更改抓取页数，是否抓取。

![](/doc/img/ca.png)


小类数据，基本Top100商品数据。

![](/doc/img/asin.png)


大类数据，很详细，包括大类排名等，可以复杂查询条件筛选，下载。

![](/doc/img/big.png)

产品趋势，你可以看到产品十几天的排名变化，价格变化。

![](/doc/img/trend.png)

导出的EXCEL

![](/doc/img/excel.png)


# 四. 欢迎咨询

如果你想咨询或学习，请发邮件或加我QQ: 459527502。

开发这个产品从2016年10月就开始了, 目前迭代从2.0开始.

此项目可以持续优化成功一个更好的平台, 因为国内目前还没有像淘宝数据参谋一样的亚马逊数据参谋. 由于高并发百万级每天导致的数据抓取速度问题, 和数据获取后的清洗和挖掘问题, 我们可以在以下方面做得更好. 

1. 首先数据抓取速度保证和爬虫部署问题, 可以采用`Docker`自动构建, 构建`kubernetes`集群进行`deployments`部署, 自动跨容和缩容爬虫服务, 分布式爬虫不再需要手工上去跑任务.
2. 其次数据保存在`MYSQL`产生的分表问题, 因为`MYSQL`是非分布式的集中式关系型数据库, 大量数据导致数据查找困难, 多表间数据`union`和`join`困难, 所以可以采用`ElasticSearch`来替换`MYSQL`, 著名的`JAVA Nutch搜索引擎框架`使用的就是`ES`.
3. 最后, 关于数据获取后的清洗和挖掘问题, 是属于离线操作问题, 保存在`ES`的数据本身支持各种搜索,`ES`的文本搜索能力超出你的想象, 一般需求可以满足, 不能满足的需求则要从`ES`抽取数据, 构建不同主题的数据仓库进行定制化挖掘. 此部分, 需要开发另外的项目.
4. 配套的`UI`网站端可以有更好的用户体验, 目前基本可以满足选款的需求, 商品的各种数据优美的显示出来.

# 免责声明

关于版权，爬虫有风险, 本人不承担由此开源项目带来的任何责任。

```
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502

	2017.7 by hunterhug
```

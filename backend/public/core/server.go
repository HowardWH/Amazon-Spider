/*
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
*/
package core

import (
	"encoding/json"
	"fmt"
	spider "github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/parrot/util"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type AmazonController struct {
	Message    string
	SpiderType string
	Port       string
}

func (c *AmazonController) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	dudu := smart2016()
	io.WriteString(rw, fmt.Sprintf(`
	<!Doctype html>
	<html>
	<head>QQ:459527502
	<meta charset="utf-8" />
	<title>超级大爬虫监控端-%s</title>
	</head>
	<body>
	<h1>%v</h1>
	SpiderType:%s<br/>Message:%s<br/>Host:%s<br/><br/>
	%s
	<div style="float:left;width:70%%";margin:40px>
	<div>
	<h1>Export URLS AGAIN</h1>
	<form action="/url" method="post">
	USER:<br/>
	<input type="text" name="user" />
	<br/>PASSWORD:<br/>
	<input type="text" name="password" />
	<input type="submit" value="RUN" />
	</form>
	</div>
	<div>
	<h1>Export IP BY YOUSERF</h1>
	<form action="/help" method="post">
	USER:<br/>
	<input type="text" name="user" />
	<br/>PASSWORD:<br/>
	<input type="text" name="password" />
	<input type="submit" value="RUN" />
	</form>
	</div>

	<div>
	<h1>Export IP DIY</h1>
	<form action="/diy" method="post">
	USER:<br/>
	<input type="text" name="user" />
	<br/>PASSWORD:<br/>
	<input type="text" name="password" />
	<br/>IPs<br/>
	<textarea name="ips" rows="20" cols="20" style="width:800px">smart@*.*.*.*:808</textarea>
	<input type="submit" value="RUN" />
	<br/>
	<br/>
	</form>
	</div>
	</div>
	<div style="float:right;width:20%%;margin:30px">
	<h1>This page you should caution!</h1>
	<img style="max-width: 100%%;" src="http://www.lenggirl.com/static/me.gif" />
	</div>
	<div>作者: 陈白痴, 版权所有: <a href="https://github.com/hunterhug">主页</a></div>
	</body>
	</html>
	`, Today, time.Now(), c.SpiderType, c.Message, c.Port, dudu))
}

func help(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		io.WriteString(rw, err.Error())
		return
	}

	c := req.Form.Get("config")
	if c == "1" {
		data, e := json.Marshal(MyConfig)
		if e == nil {
			io.WriteString(rw, string(data))
			return
		}
	} else if c == "2" {
		os.Exit(0)
	}
	user := req.Form.Get("user")
	password := req.Form.Get("password")
	if user == "jinhan" && password == "459527502" {
		io.WriteString(rw, Sentiptoredis(IPPOOL))
	} else {
		io.WriteString(rw, "not allow!!")
	}
}

func url(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		io.WriteString(rw, err.Error())
		return
	}
	user := req.Form.Get("user")
	password := req.Form.Get("password")
	if user == "jinhan" && password == "459527502" {
		result, err := BasicDb.Select(MyConfig.Urlsql)
		if err != nil {
			io.WriteString(rw, err.Error())
			return
		}
		urls := []string{}
		for _, index := range result {
			urls = append(urls, index["id"].(string)+"|"+index["url"].(string)+"|"+index["name"].(string)+"|"+index["bigpname"].(string)+"|"+index["page"].(string))
		}
		s := "total:" + util.IS(len(urls)) + " urls\n"
		for _, url := range urls {
			_, err := RedisClient.Lpush(MyConfig.Urlpool, url)
			if err != nil {
				s = s + fmt.Sprintf("error:%v,%v\n", url, err)
			}
		}
		io.WriteString(rw, s)
	} else {
		io.WriteString(rw, "not allow!!")
	}
}

type mixx struct {
	Url    string  `json:"url"`
	Result []mixxx `json:"result"`
}
type mixxx struct {
	Ip   string `json:"ip:port"`   //"ip:port": "67.207.95.138:8080",
	Type string `json:"http_type"` //"http_type": "HTTPS",
	An   string `json:"anonymous"` //"anonymous": "高匿",
	Isp  string `json:"isp"`       //"isp": "null",
	C    string `json:"country"`   //"country": "美国"
}

// http://127.0.0.1:12345/mi?user=jinhan&password=459527502
func mi(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		io.WriteString(rw, err.Error())
		return
	}
	user := req.Form.Get("user")
	password := req.Form.Get("password")
	orderid := req.Form.Get("orderid")
	if user == "jinhan" && password == "459527502" {
		if orderid == "" {
			return
		}
		num, e := RedisClient.Llen(MyConfig.Proxypool)
		if e == nil && num > 5 {
			io.WriteString(rw, fmt.Sprintf("still has ip:%d", num))
			return
		}
		url := "http://proxy.mimvp.com/api/fetch.php?orderid=%s&num=100&result_format=json&anonymous=5&result_fields=1,2,3,4,5&http_type=1,2,5"
		sp := spider.NewAPI()
		sp.Url = fmt.Sprintf(url, orderid)
		data, err := sp.Get()
		if err != nil {
			io.WriteString(rw, err.Error())
			return
		}
		r := new(mixx)
		err = json.Unmarshal(data, r)
		if err != nil {
			io.WriteString(rw, err.Error())
			return
		}
		if len(r.Result) == 0 {
			io.WriteString(rw, "zero")
			return
		}
		ipsmart2016 := []string{}
		for _, i := range r.Result {
			if i.Type == "Socks5" {
				i.Ip = "socks5://" + i.Ip
			} else {
				i.Ip = "http://" + i.Ip
			}
			ipsmart2016 = append(ipsmart2016, i.Ip)
		}
		io.WriteString(rw, Sentiptoredis(ipsmart2016))
	} else {
		io.WriteString(rw, "not allow!!")
	}
}
func diy(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		io.WriteString(rw, err.Error())
		return
	}
	user := req.Form.Get("user")
	password := req.Form.Get("password")
	if user == "jinhan" && password == "459527502" {
		ipsmart2016 := []string{}
		ipstring := req.Form.Get("ips")
		tempips := strings.Split(ipstring, "\n")
		for _, tempip := range tempips {
			ip := strings.TrimSpace(strings.Replace(tempip, "\r", "", -1))
			dudu := strings.Split(ip, ".")
			if len(dudu) != 4 {
				continue
			} else {
				IPdudu := true
				for _, d := range dudu {
					tempd := d
					d1 := strings.Split(d, "@")
					if len(d1) == 2 {
						tempd = d1[1]
					}
					if len(d1) > 2 {
						IPdudu = false
						break
					}
					d2 := strings.Split(tempd, ":")
					if len(d2) > 2 {
						IPdudu = false
						break
					}
					tempd = d2[0]
					dnum, de := util.SI(tempd)
					if de != nil {
						IPdudu = false
						break
					}
					if dnum > 255 || dnum <= 0 {
						IPdudu = false
						break
					}
				}
				if IPdudu {
					ipsmart2016 = append(ipsmart2016, ip)
				}
			}
		}

		io.WriteString(rw, Sentiptoredis(ipsmart2016))
	} else {
		io.WriteString(rw, "not allow!!")
	}
}
func ServePort(host string, ac *AmazonController) error {
	//:8080
	ac.Port = host
	http.Handle("/", ac)
	http.HandleFunc("/help", help)
	http.HandleFunc("/diy", diy)
	http.HandleFunc("/url", url)
	// http://proxy.mimvp.com/api/fetch.php?orderid=860170716115639588&num=100&result_format=json&anonymous=5&result_fields=1,2,3,4,5&http_type=1,2,5
	http.HandleFunc("/mi", mi)
	err := http.ListenAndServe(host, nil)
	return err
}

/*
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
*/

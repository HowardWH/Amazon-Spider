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
	"math/rand"
	"strings"
	"sync"

	spider "github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/parrot/util"
)

var (
	Spiders = &_Spider{brower: make(map[string]*spider.Worker)}
	Ua      = map[int]string{}
)

type _Spider struct {
	mux    sync.RWMutex
	brower map[string]*spider.Worker
}

func (sb *_Spider) Get(name string) (b *spider.Worker, ok bool) {
	sb.mux.RLock()
	b, ok = sb.brower[name]
	sb.mux.RUnlock()
	return
}

func (sb *_Spider) Set(name string, b *spider.Worker) {
	sb.mux.Lock()
	sb.brower[name] = b
	sb.mux.Unlock()
	return
}

func (sb *_Spider) Delete(name string) {
	sb.mux.Lock()
	delete(sb.brower, name)
	sb.mux.Unlock()
	return
}
func init() {
	Ua[0] = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36"
	utxt := Dir + "/config/ua.txt"
	temp, err := util.ReadfromFile(utxt)
	if err != nil {
		panic(err.Error())
	} else {
		uas := strings.Split(string(temp), "\n")

		for i, ua := range uas {
			Ua[i] = strings.TrimSpace(strings.Replace(ua, "\r", "", -1))
		}
	}

}
func Download(ip string, url string) ([]byte, error) {
	if strings.Contains(ip, "*") {
		return NonProxyDownload(ip, url)
	}
	browser, ok := Spiders.Get(ip)
	if ok {
		browser.Url = url
		content, err := browser.Get()
		spider.Logger.Debugf("url:%s,status:%d,ip:%s,ua:%s", url, browser.UrlStatuscode, ip, browser.Header.Get("User-Agent"))
		return content, err
	} else {
		browser, _ := spider.New(ip)
		browser.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		browser.Header.Set("Accept-Language", "en-US;q=0.8,en;q=0.5")
		browser.Header.Set("Connection", "keep-alive")
		if strings.Contains(url, "www.amazon.co.jp") {
			browser.Header.Set("Host", "www.amazon.co.jp")
		} else if strings.Contains(url, "www.amazon.de") {
			browser.Header.Set("Host", "www.amazon.de")
		} else if strings.Contains(url, "www.amazon.co.uk") {
			browser.Header.Set("Host", "www.amazon.co.uk")
		} else {
			browser.Header.Set("Host", "www.amazon.com")
		}
		browser.Header.Set("Upgrade-Insecure-Requests", "1")
		browser.Header.Set("User-Agent", Ua[rand.Intn(len(Ua)-1)])
		browser.Url = url
		Spiders.Set(ip, browser)
		content, err := browser.Get()
		spider.Logger.Debugf("url:%s,status:%d,ip:%s,ua:%s", url, browser.UrlStatuscode, ip, browser.Header.Get("User-Agent"))
		return content, err
	}
}

func NonProxyDownload(ip string, url string) ([]byte, error) {
	browser, ok := Spiders.Get(ip)
	if ok {
		browser.Url = url
		content, err := browser.Get()
		spider.Logger.Debugf("url:%s,status:%d,ip:%s,ua:%s", url, browser.UrlStatuscode, ip, browser.Header.Get("User-Agent"))
		return content, err
	} else {
		browser, _ := spider.New(nil)
		browser.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		browser.Header.Set("Accept-Language", "en-US;q=0.8,en;q=0.5")
		browser.Header.Set("Connection", "keep-alive")
		if strings.Contains(url, "www.amazon.co.jp") {
			browser.Header.Set("Host", "www.amazon.co.jp")
		} else if strings.Contains(url, "www.amazon.de") {
			browser.Header.Set("Host", "www.amazon.de")
		} else if strings.Contains(url, "www.amazon.co.uk") {
			browser.Header.Set("Host", "www.amazon.co.uk")
		} else {
			browser.Header.Set("Host", "www.amazon.com")
		}
		browser.Header.Set("Upgrade-Insecure-Requests", "1")
		browser.Header.Set("User-Agent", Ua[rand.Intn(len(Ua)-1)])
		browser.Url = url
		Spiders.Set(ip, browser)
		content, err := browser.Get()
		spider.Logger.Debugf("url:%s,status:%d,ip:%s,ua:%s", url, browser.UrlStatuscode, ip, browser.Header.Get("User-Agent"))
		return content, err
	}
}

func GetIP() string {
	spider.Logger.Debug("Get IP...")
	iptemp, ierr := RedisClient.Brpop(0, MyConfig.Proxypool)
	// ip null return,maybe forever not happen
	if ierr != nil {
		panic("ip:" + ierr.Error())
	}
	ip := iptemp[1]
	spider.Logger.Debug("Get IP done:" + ip)
	return ip
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

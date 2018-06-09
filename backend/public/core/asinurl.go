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
	"errors"
	"github.com/hunterhug/parrot/util"
	"strings"
)

// ip and url download,if error from zero again!
// ip := "104.128.124.122:808"
// url := "https://www.amazon.com/dp/B01MTJ1E4Q"
// filename:B01MTJ1E4Q
func GetAsinUrl(ip string, url string) ([]byte, error) {
	filename := strings.Split(url, "/dp/")
	if len(filename) != 2 {
		return nil, errors.New("404")
	}
	keepdirtemp := MyConfig.Datadir + "/asin/" + Today + "/" + filename[1] + ".html"
	if MyConfig.Asinlocalkeep {
		if util.FileExist(keepdirtemp) {
			AmazonAsinLog.Debugf("FileExist:%s", keepdirtemp)
			return util.ReadfromFile(keepdirtemp)
		}
		if util.FileExist(keepdirtemp + "sql") {
			AmazonAsinLog.Debugf("FileExist: % sql", keepdirtemp)
			return util.ReadfromFile(keepdirtemp + "sql")
		}
	}
	content, err := Download(ip, url)
	if err != nil {
		return nil, err
	}
	if IsRobot(content) {
		return nil, errors.New("robot")
	}
	if Is404(content) {
		return nil, errors.New("404")
	}
	if MyConfig.Asinlocalkeep {
		util.SaveToFile(keepdirtemp, content)
	}
	return content, nil

}

// most import
func GetAsinUrls() error {
	AmazonAsinLog.Log("Start Get Asin url")
	ip := GetIP()

	// before use, send to hash pool
	ipbegintimes := util.GetSecond2DateTimes(util.GetSecondTimes())
	RedisClient.Hset(MyConfig.Proxyhashpool, ip, ipbegintimes)

	// do a lot url still can't pop url
	for {
		// take url!block!!!
		// url such like https://www.amazon.com/dp/B01MTJ1E4Q
		// take a url and throw it into deal pool
		url, err := RedisClient.Brpoplpush(MyConfig.Asinpool, MyConfig.Asindealpool, 0)
		if err != nil {
			return err
		}
		exist, _ := RedisClient.Hexists(MyConfig.Asinhashpool, url)
		if exist {
			AmazonAsinLog.Errorf("exist %s", url)
			continue
		}

		urlbegintime := util.GetSecond2DateTimes(util.GetSecondTimes())

		content := []byte("")
		err = nil
		// when error loop
		for {
			content, err = GetAsinUrl(ip, url)
			spider, ok := Spiders.Get(ip)
			if err == nil {
				break
			} else {
				if strings.Contains(err.Error(), "404") {
					break
				}
				if strings.Contains(err.Error(), "robot") {
					if ok {
						spider.Errortimes = spider.Errortimes + 1
					}
				}
				if ok {
					AmazonAsinLog.Errorf("get %s fail(%d),total(%d) error:%s,ip:%s", url, spider.Errortimes, spider.Fetchtimes, err.Error(), ip)
				}
			}
			// if proxy ip err more than config, change ip
			if ok && spider.Errortimes > MyConfig.Proxymaxtrytimes {
				// die sent
				ipendtimes := util.GetSecond2DateTimes(util.GetSecondTimes())
				insidetemp := ipbegintimes + "|" + ipendtimes + "|" + util.IS(spider.Fetchtimes-spider.Errortimes) + "|" + util.IS(spider.Errortimes)
				RedisClient.Hset(MyConfig.Proxyhashpool, ip, insidetemp)
				// you know it
				Spiders.Delete(ip)
				// get new proxy again
				ip = GetIP()
				ipbegintimes = util.GetSecond2DateTimes(util.GetSecondTimes())
				RedisClient.Hset(MyConfig.Proxyhashpool, ip, ipbegintimes)
			}
		}
		if err != nil && strings.Contains(err.Error(), "404") {
			// 404 set asin invaild
			err = SetAsinInvalid(url)
			if err != nil {
				AmazonAsinLog.Errorf("%s set invalid error:%s", url, err.Error())
			}
		} else {
			// parse detail
			info := ParseDetail(url, content)
			// insert, error you still ignore
			err := InsertDetailMysql(info)
			if err != nil {
				AmazonAsinLog.Errorf("%s mysql error:%s", url, err.Error())
			}
		}
		// done! rem redis deal pool
		RedisClient.Lrem(MyConfig.Asindealpool, 0, url)
		// throw it to a hash pool
		urlendtimes := util.GetSecond2DateTimes(util.GetSecondTimes())
		RedisClient.Hset(MyConfig.Asinhashpool, url, urlbegintime+"|"+urlendtimes)
	}
	return nil
}

func GetNoneProxyAsinUrls(taskname string) error {
	AmazonAsinLog.Log("Start Get Asin url")
	ip := "*" + taskname

	// do a lot url still can't pop url
	for {
		// take url!block!!!
		// url such like https://www.amazon.com/dp/B01MTJ1E4Q
		// take a url and throw it into deal pool
		url, err := RedisClient.Brpoplpush(MyConfig.Asinpool, MyConfig.Asindealpool, 0)
		if err != nil {
			return err
		}

		exist, _ := RedisClient.Hexists(MyConfig.Asinhashpool, url)
		if exist {
			AmazonAsinLog.Errorf("exist %s", url)
			continue
		}
		urlbegintime := util.GetSecond2DateTimes(util.GetSecondTimes())

		content := []byte("")
		err = nil
		// when error loop
		for {
			content, err = GetAsinUrl(ip, url)
			spider, ok := Spiders.Get(ip)
			if err == nil {
				break
			} else {
				if strings.Contains(err.Error(), "404") {
					break
				}
				if strings.Contains(err.Error(), "robot") {
					if ok {
						spider.Errortimes = spider.Errortimes + 1
					}
				}
				if ok {
					AmazonAsinLog.Errorf("get %s fail(%d),total(%d) error:%s,ip:%s", url, spider.Errortimes, spider.Fetchtimes, err.Error(), ip)
				}
			}
			// if proxy ip err more than config, change ip
			if ok && spider.Errortimes > MyConfig.Proxymaxtrytimes {
				// you know it
				Spiders.Delete(ip)
			}
		}
		if err != nil && strings.Contains(err.Error(), "404") {
			// 404 set asin invaild
			err = SetAsinInvalid(url)
			if err != nil {
				AmazonAsinLog.Errorf("%s set invalid error:%s", url, err.Error())
			}
		} else {
			// parse detail
			info := ParseDetail(url, content)
			// insert, error you still ignore
			err := InsertDetailMysql(info)
			if err != nil {
				AmazonAsinLog.Errorf("%s mysql error:%s", url, err.Error())
			} else {
				AmazonAsinLog.Debug("Insert!!")
			}
		}
		// done! rem redis deal pool
		RedisClient.Lrem(MyConfig.Asindealpool, 0, url)
		// throw it to a hash pool
		urlendtimes := util.GetSecond2DateTimes(util.GetSecondTimes())
		RedisClient.Hset(MyConfig.Asinhashpool, url, urlbegintime+"|"+urlendtimes)
	}
	return nil
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

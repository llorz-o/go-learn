package reptile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var reEmail = `\w+@\w+\.\w+`
var reIdCard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
var reLink = `href="(https?://[\s\S]+?)"`
var rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
var reImg    = `(https?:)?\/\/[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`


func getPageStr(url string) string {
	resp, err := http.Get(url)
	handleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	create, err := os.Create("reptile.html")
	handleError(err,"os.Create")
	create.Write(pageBytes)
	handleError(err, "ioutil.ReadAll")
	return string(pageBytes)
}

func reptile(reg string,url string) {
	for _, result := range regexp.MustCompile(reg).FindAllStringSubmatch(getPageStr(url), -1) {
		fmt.Println(result)
	}
}

func FuncReptile2() {
	//reptile(reEmail,"https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	//reptile(reIdCard,"https://henan.qq.com/a/20171107/069413.htm")
	//reptile(reLink,"http://www.baidu.com/s?wd=%E8%B4%B4%E5%90%A7%20%E7%95%99%E4%B8%8B%E9%82%AE%E7%AE%B1&rsv_spt=1&rsv_iqid=0x98ace53400003985&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_dl=ib&rsv_sug2=0&inputT=5197&rsv_sug4=6345")
	//reptile(rePhone,"https://www.zhaohaowang.com/") // 链接不可用
	reptile(reImg,"http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")
}

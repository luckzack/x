package geo

import (
	"log"
	"strings"

	"github.com/wangtuanjie/ip17mon"
)

func Init(dat string) {
	if len(dat) == 0 {
		dat = "./var/mydata4vipday2-4.dat"
	}
	if err := ip17mon.Init(dat); err != nil {
		log.Fatal("get the isp fail:%v", err)
	}
}

func GetLocation(ip string) (*ip17mon.LocationInfo, error) {
	loc, err := ip17mon.Find(ip)
	if err != nil {
		return nil, err
	}

	if strings.Contains(loc.Isp, "电信") {
		loc.Isp = "电信"
	} else if strings.Contains(loc.Isp, "联通") {
		loc.Isp = "联通"
	} else if strings.Contains(loc.Isp, "移动") {
		loc.Isp = "移动"
	}

	//loc.Isp = consts.ClassifyMyISP(loc.Isp)

	return loc, nil

}

package baidu_netdisk_sharelink_check

import (
	"strings"

	"github.com/irebit/baidu_netdisk_sharelink_check/request"
)

const DestURLPrefix = "https://pan.baidu.com/share/init?surl="

type ShareLink string

// func (l ShareLink) ToString() string {
// 	return fmt.Sprintf("%v", l)
// }

func (l ShareLink) Check() (bool, error) {
	_, resp, err := request.Get(string(l))
	if len(resp.Header["Location"]) == 1 {
		if strings.Contains(resp.Header["Location"][0], DestURLPrefix) {
			return true, err
		}
	}
	return false, err
}

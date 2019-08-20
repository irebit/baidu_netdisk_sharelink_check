package baidu_netdisk_sharelink_check

import (
	"log"
	"testing"
)

func TestCheck(t *testing.T) {
	//suc
	var link ShareLink = "https://pan.baidu.com/s/1D_vNhmirNjGC3gkdxwGPYw"
	log.Println(link.Check())
	//fail
	var link2 ShareLink = "https://pan.baidu.com/s/1bQRjzr8x7XvKj8L3_iv9xw"
	log.Println(link2.Check())
}

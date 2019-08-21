package baidu_netdisk_sharelink_check

import (
	"log"
	"testing"
)

func TestCheck(t *testing.T) {
	//suc
	log.Println(ShareLink("https://pan.baidu.com/s/1D_vNhmirNjGC3gkdxwGPYw").Check())
	//fail
	log.Println(ShareLink("https://pan.baidu.com/s/1bQRjzr8x7XvKj8L3_iv9xw").Check())
}

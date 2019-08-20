# baidu_netdisk_sharelink_check
检查百度网盘分享链接是否可用

## 实现方法
利用正常分享链接会302跳转到目标页，从而判断是否为正常页面

## 如何调用

``` golang
import (
    "log"
    "github.com/irebit/baidu_netdisk_sharelink_check"
)

func main() {
	//suc
	var link baidu_netdisk_sharelink_check.ShareLink = "https://pan.baidu.com/s/1D_vNhmirNjGC3gkdxwGPYw"
	log.Println(link.Check())
	//fail
	var link2 baidu_netdisk_sharelink_check.ShareLink = "https://pan.baidu.com/s/1bQRjzr8x7XvKj8L3_iv9xw"
	log.Println(link2.Check())
}
```
package main

import (
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func system() string {
	sysType := runtime.GOOS
	if sysType == "linux" {
		return "bash"
	}

	if sysType == "windows" {
		return "cmd"
	}
	return ""
}

func IP() string {
	//待获取的网页数据
	url := "http://members.3322.org/dyndns/getip"
	// 根据URL获取资源
	res, _ := http.Get(url)
	// 读取资源数据 body: []byte
	body, _ := ioutil.ReadAll(res.Body)
	// 关闭资源流
	res.Body.Close()
	IP := strings.Replace(string(body), "\n", "", -1)
	//返回IP地址
	return string(IP)
}
func SSH(message string) string {
	if message == "打开" || message == "开启" || message == "启动" {
		err := Shell("/usr/syno/sbin/synoservice --start ssh-shell", true)
		if len(err) == 0 {
			return message + "ssh 成功!!!"
		}

	}

	if message == "停止" || message == "关闭" || message == "终止" {
		err := Shell("/usr/syno/sbin/synoservice --stop ssh-shell", true)
		if len(err) > 0 {
			return err
		}
	}

	if message == "重启" {
		err := Shell("/usr/syno/sbin/synoservice --restart ssh-shell", true)
		if len(err) == 0 {
			return message + "ssh 成功!!!"
		}
	}
	return ""
}

func Shell(cmd string, shell bool) string {
	if shell {
		if runtime.GOOS == "linux" {
			out, err := exec.Command(system(), "-c", cmd).Output()
			if err != nil {
				panic("some error found")
			}
			return string(out)
		} else {
			out, err := exec.Command(system(), "/C", cmd).Output()
			if err != nil {
				panic("some error found")
			}
			output, _ := simplifiedchinese.GBK.NewDecoder().Bytes(out)
			return string(output)
		}
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return string(out)
	}
	return ""
}

func handle(message string) string {
	MESSAGE := strings.ToUpper(message)
	if MESSAGE == "IP" {
		return "公网IP：" + IP()
	}

	if MESSAGE == "NAS" || MESSAGE == "群晖" {
		return "群晖IP：" + IP() + ":5000"
	}

	if MESSAGE == "OP" || MESSAGE == "OPENWRT" || MESSAGE == "软路由" {
		return "软路由：" + IP() + ":8"
	}

	if MESSAGE == "ASUS" || MESSAGE == "路由器" {
		return "路由器：" + IP() + ":6"
	}

	if MESSAGE == "ALIST" {
		return "网盘管理：" + IP() + ":5244"
	}

	if MESSAGE == "VS" {
		return "VS_Code：" + IP() + ":8081"
	}

	if MESSAGE[6:] == "SSH" {
		return SSH(MESSAGE[0:6])
	}

	if MESSAGE[0:6] == "命令" {
		//println(Shell(message[6:], true))
		return Shell(message[6:], true)

	}

	return ""
}

func POST(ctx *gin.Context) {
	dataReader := ctx.Request.Body
	rawData, _ := ioutil.ReadAll(dataReader)
	postType := gjson.Get(string(rawData), "post_type").String()
	if postType == "message" {
		message := gjson.Get(string(rawData), "message").String()
		println(message)
		if len(message) > 0 {
			params := handle(message)
			if len(params) > 0 {
				data := gin.H{"reply": params}
				ctx.JSON(http.StatusOK, data)
			}
		}
	}
}

func GET(ctx *gin.Context) {
	data := gin.H{"POST": "OK"}
	ctx.JSON(http.StatusOK, data)
}

func main() {

	app := gin.Default()
	app.GET("/", GET)
	app.POST("/api", POST)
	println(app)
	app.Run(":5701")

}

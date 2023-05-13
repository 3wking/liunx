package main

/*/小Q助手/*/

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/denverdino/aliyungo/dns"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func DDNS() string {
	// 读取配置文件信息
	AccessKeyID := "LTAITH5SfzDHXM9T"
	AccessKeySecret := "ZdLWBCDBExC5gKnS4L4cCZpiT9Z7dm"
	DomainName := "3wking.com"
	RR := "@"
	// 获取公网IP信息
	publicIP := IP()
	fmt.Println("当前公网IP：" + publicIP)
	// 连接阿里云服务器，获取DNS信息
	client := dns.NewClient(AccessKeyID, AccessKeySecret)
	client.SetDebug(false)
	domainInfo := new(dns.DescribeDomainRecordsArgs)
	domainInfo.DomainName = DomainName
	oldRecord, err := client.DescribeDomainRecords(domainInfo)
	if err != nil {
		return "链接错误"
	}

	var exsitRecordID string
	for _, record := range oldRecord.DomainRecords.Record {
		if record.DomainName == DomainName && record.RR == RR {
			if record.Value == publicIP {
				return "域名IP与公网IP相同,无须解析。"
			}
			exsitRecordID = record.RecordId
		}
	}

	if 0 < len(exsitRecordID) {
		// 有配置记录，则匹配配置文件，进行更新操作
		updateRecord := new(dns.UpdateDomainRecordArgs)
		updateRecord.RecordId = exsitRecordID
		updateRecord.RR = RR
		updateRecord.Value = publicIP
		updateRecord.Type = dns.ARecord
		rsp := new(dns.UpdateDomainRecordResponse)
		rsp, err := client.UpdateDomainRecord(updateRecord)
		if nil != err {
			return "修改解析失败"
			println("修改解析失败", err)
		} else {
			return "修改解析成功"
			println("修改解析成功", rsp)
		}
	} else {
		// 没有找到配置记录，那么就新增一个
		newRecord := new(dns.AddDomainRecordArgs)
		newRecord.DomainName = DomainName
		newRecord.RR = RR
		newRecord.Value = publicIP
		newRecord.Type = dns.ARecord
		rsp := new(dns.AddDomainRecordResponse)
		rsp, err = client.AddDomainRecord(newRecord)
		if nil != err {
			return "添加DNS解析失败"
			println("修改解析失败", err)

		} else {
			return "添加DNS解析成功"
			println("修改解析成功", rsp)
		}
	}
	return ""
}

func system() string {
	sysType := runtime.GOOS
	if sysType == "linux" {
		return "bash"
	}

	if sysType == "windows" {
		return "PowerShell"
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
		Shell := exec.Command(system(), "-c", cmd)
		if runtime.GOOS == "windows" {
			Shell = exec.Command(system(), "/C", cmd)
		}
		out, err := Shell.CombinedOutput()
		if err != nil {
			//panic("some error found")
			println("some error found")
		}
		if runtime.GOOS == "windows" {
			output, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(out))
			return string(output)
		}
		return string(out)

	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			//panic("some error found")
			println("some error found")
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
		return "群晖：" + IP() + ":5000"
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
		return "VS编辑器：" + IP() + ":8081"
	}

	if len(MESSAGE) >= 6 {
		if MESSAGE[6:] == "SSH" {
			return SSH(MESSAGE[0:6])
		}

		if MESSAGE[0:6] == "命令" || MESSAGE[0:6] == "运行" {
			return Shell(message[6:], true)
		}

		if MESSAGE[0:6] == "解析" {
			return DDNS()
		}

	}
	return ""
}

func POST(c *gin.Context) {
	dataReader := c.Request.Body
	rawData, _ := ioutil.ReadAll(dataReader)
	postType := gjson.Get(string(rawData), "post_type").String()
	if postType == "message" {
		message := gjson.Get(string(rawData), "message").String()
		println(message)
		if len(message) > 0 {
			params := handle(message)
			if len(params) > 0 {
				data := gin.H{"reply": params}
				c.JSON(http.StatusOK, data)
			}
		}
	}
}

func GET(c *gin.Context) {
	//data := gin.H{"POST": "OK"}
	//c.JSON(http.StatusOK, data)
	c.String(200, "OK")
}

func main() {
	app := gin.Default()
	app.GET("/", GET)
	app.POST("/api", POST)
	app.Run(":5701")
}

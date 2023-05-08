# uduntu
## iso镜像下载 ：
* [官网](https://ubuntu.com/#download)

## Ubuntu20.04软件源更换：
##### 备份源：
```sh
sudo cp /etc/apt/sources.list /etc/apt/sources.list.bak		#备份
```
```sh
sudo mv /etc/apt/sources.list.bak /etc/apt/sources.list		#还原
```
##### 阿里：
```sh
echo "#添加阿里源" > /etc/apt/sources.list
echo "deb http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse" >> /etc/apt/sources.list
echo "deb-src http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse" >> /etc/apt/sources.list
echo "deb http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse" >> /etc/apt/sources.list
echo "deb-src http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse" >> /etc/apt/sources.list
echo "deb http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse" >> /etc/apt/sources.list
echo "deb-src http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse" >> /etc/apt/sources.list
echo "deb http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse" >> /etc/apt/sources.list
echo "deb-src http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse" >> /etc/apt/sources.list
echo "deb http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse" >> /etc/apt/sources.list
echo "deb-src http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse" >> /etc/apt/sources.list
```
##### #添加清华源：
```sh
echo "#添加清华源" > /etc/apt/sources.list
echo "deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal main restricted universe multiverse" >> /etc/apt/sources.list
echo "# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal main restricted universe multiverse" >> /etc/apt/sources.list
echo "deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates main restricted universe multiverse" >> /etc/apt/sources.list
echo "# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-updates main restricted universe multiverse" >> /etc/apt/sources.list
echo "deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-backports main restricted universe multiverse" >> /etc/apt/sources.list
echo "# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-backports main restricted universe multiverse" >> /etc/apt/sources.list
echo "deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security main restricted universe multiverse" >> /etc/apt/sources.list
echo "# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ focal-security main restricted universe multiverse multiverse" >> /etc/apt/sources.list
```


## 软件源：
##### 添加PPA源到Source list中：
```sh
sudo apt-add-repository  ppa_source_name
```
##### apt-add-repository：
```sh
apt-get install software-properties-common
```
```sh
apt-get install python-software-propertiessudo
```
##### 更新证书：
```sh
sudo apt update -y
sudo apt-get install apt-transport-https ca-certificates -y
sudo update-ca-certificates
```
## 安装deb软件包
##### apt命令：
```sh
sudo apt install path_to_deb_file
```
##### dpkg命令：
```sh
sudo dpkg -i path_to_deb_file
```
##### 赖项的错误：
```sh
sudo apt install -f
```
## 查看deb软件包
##### apt命令：
```sh
sudo apt list --installed
```
##### dpkg命令：
```sh
dpkg -l
```

## 移除deb软件包
##### apt命令：
```sh
sudo apt remove program_name
```
##### dpkg命令：
```sh
dpkg -r program_name
```
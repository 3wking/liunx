# uduntu
## iso镜像下载 ：
* [官网](https://ubuntu.com/#download)

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
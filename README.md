# linux
## linux禁用swap
#### 不重启电脑，禁用启用swap，立刻生效
###### 禁用命令:
```sh
sudo swapoff -a
```
###### 启用命令:
```sh
sudo swapon -a
```
###### 查看交换分区的状态:
```sh
sudo free -m
```
## 时间
#### 最接近的时间,四舍五入
###### 5分钟
```
echo $(date -d @$(( (($(date +%s) + 150) / 300) * 300)) "+%Y_%m_%d_%H_%M")
```
###### 2分钟
```
echo $(date -d @$(( (($(date +%s) + 150) / 120) * 120)) "+%Y_%m_%d_%H_%M")
```

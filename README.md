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
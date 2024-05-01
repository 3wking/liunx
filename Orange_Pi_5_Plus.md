# 5_Plus
## LED控制
#### 设置绿灯LED
###### 先进入绿灯的设置目录
```
cd /sys/class/leds/green_led
```
###### 设置绿灯停止闪烁
```
echo none > trigger
```
###### 设置绿灯常亮
```
echo default-on > trigger
```
###### 设置绿灯闪烁
```
 echo heartbeat > trigger
```
#### 设置蓝灯LED
###### 先进入蓝灯的设置目录
```
cd /sys/class/leds//blue_led
```
###### 设置蓝灯停止闪烁
```
echo none > trigger
```
###### 设置蓝灯常亮
```
echo default-on > trigger
```
###### 设置蓝灯闪烁
```
 echo heartbeat > trigger
```
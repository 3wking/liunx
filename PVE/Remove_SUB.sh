#!/bin/bash
RED_COLOR='\e[1;31m' #红色
GREEN_COLOR='\e[1;32m' #绿色
RES='\e[0m' #尾
#安装
function Install() {
	dir="/usr/share/javascript/proxmox-widget-toolkit/proxmoxlib.js"
	rows=$(($(grep -n "No valid subscription" $dir |awk -F ':' '{print $1}')-1))
	text=`sed -n ${rows}p $dir | grep "Ext.Msg.show({"`
	if [[ $text ]]; then
		echo -e "\r\n${GREEN_COLOR}寻找结果为：${#text}${RES}"
		echo -e "${GREEN_COLOR}当前内容为：`sed -n ${rows}p $dir`${RES}"
		sed -i ''$rows's/show/noshow/g' $dir >/dev/null 2>&1 
		if [ $? -ne 0 ]; then
			echo -e "${RED_COLOR}内容修改失败.${RES}"
			echo -e "${RED_COLOR}当前内容为： `sed -n ${rows}p $dir`${RES}\r\n"
			exit 1
		fi
		echo -e "${GREEN_COLOR}内容更改成功...${RES}"
		echo -e "${GREEN_COLOR}当前内容为：`sed -n ${rows}p $dir`${RES}\r\n"
	else
		echo -e "\r\n${RED_COLOR}寻找结果为：${#text}${RES}"
		echo -e "${RED_COLOR}内容更改失败.${RES}"
		echo -e "${RED_COLOR}当前内容为： `sed -n ${rows}p $dir`${RES}\r\n"
	fi
}
Install
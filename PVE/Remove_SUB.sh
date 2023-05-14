#!/bin/bash
RED_COLOR='\e[1;31m' #红色
GREEN_COLOR='\e[1;32m' #绿色
RES='\e[0m' #尾
#安装
function Install() {
	dir="/usr/share/javascript/proxmox-widget-toolkit/proxmoxlib.js"
	rows=$(($(grep -n "No valid subscription" $dir |awk -F ':' '{print $1}')-1))
	text=`sed -n ${tmp}p $dir | grep "Ext.Msg.show({"`
	if [ $text ]; then
		sed -i ''$tmp's/show/noshow/g' $dir
		echo -e "${GREEN_COLOR}内容更改成功...${RES}"
	else
		echo -e "${RED_COLOR}内容更改失败.${RES}"

	fi
}
Install
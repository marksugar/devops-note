#!/bin/bash
spath=/data/cmdb_workspace/export
dpath=/data/wwwroot
gitVersion=${gitVersion:-"git pull -u origin master"}
logfile="/data/script/upcodeOneByMany.log"

m_php_config(){
	for i in $1;do
		if [ ! -d $spath/${i}vue_config ];then echo "无效的web id"; continue; fi
		files1=$(ls $spath | grep ${i}vue_config)
		cd $spath/$files1 && if [ "${3}" == "" ]; then $gitVersion; else $gitVersion && git checkout ${3}; fi
		cp -r $spath/$files1/wcphpsec/config.php $dpath/${i}vue/m/php/
		if [ $? -ne 0 ]; then
			echo "failed $i" >> $logfile
		else
			echo "succeed $i" >> $logfile
		fi
		done
}
vue_wap(){
	for i in $1;do
		if [ ! -d $spath/${i}vue_wap ];then echo "无效的web id"; continue; fi
		files1=$(ls $spath | grep ${i}vue_wap)
		cd $spath/$files1 &&  if [ "${3}" == "" ]; then $gitVersion; else $gitVersion && git checkout ${3}; fi
		cp -r $spath/$files1/* $dpath/${i}vue/
		if [ $? -ne 0 ]; then
			echo "failed $i" >> $logfile
		else
			echo "succeed $i" >> $logfile
		fi
		done
}
vue_pc(){
	for i in $1;do
		if [ ! -d $spath/${i}vue_pc ];then echo "无效的web id"; continue; fi
		files1=$(ls $spath | grep ${i}vue_pc)
		cd $spath/$files1 && if [ "${3}" == "" ]; then $gitVersion; else $gitVersion && git checkout ${3}; fi
		cp -r $spath/$files1/* $dpath/${i}vue/
		if [ $? -ne 0 ]; then
			echo "failed $i" >> $logfile
		else
			echo "succeed $i" >> $logfile
		fi
		done
}
m_php(){
	for i in $1;do
		if [ ! -d $spath/${i}vue_php ];then echo "无效的web id"; continue; fi
		files1=$(ls $spath | grep ${i}vue_php)
		cd $spath/$files1 &&  if [ "${3}" == "" ]; then $gitVersion; else $gitVersion && git checkout ${3}; fi
		cp -r $spath/$files1/* $dpath/${i}vue/
		files2=$(ls $spath | grep ${i}vue_config)
		cp -r $spath/$files2/wcphpsec/config.php $dpath/${i}vue/m/php/
		if [ $? -ne 0 ]; then
			echo "failed $i" >> $logfile
		else
			echo "succeed $i" >> $logfile
		fi
		done
}

ARGS=1
if [ $# -ne "$ARGS" ]; then
    echo "Please input one arguement:"
fi
case $1 in
    m_php_config)
		result=`m_php_config "$2" $3`
		echo -e "$result\n"
		;;
    m_php)
		result=`m_php "$2" $3`
		echo -e "$result\n"
		;;
    vue_pc)
		result=`vue_pc "$2" $3`
		echo -e "$result\n"
		;;
    vue_wap)
		result=`vue_wap "$2" $3`
		echo -e "$result\n"
		;;
		*)
	echo "Usage:$0(m_php_config|m_php|vue_pc|vue_wap)"
		;;
esac
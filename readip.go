package main

import (
	"bufio"
	"os"
)

var srcIP []string
/*********************************************************************
 * Function: 打开IP白名单 将白名单中的ip读取出来
 * Description:
 *	syntax --  readIPFile() (src []string, err error)
 *
 * 	input :
 * 	return value :
 *						src []string: 存放IP的字符串数组
 *						err 错误信息
 * Designer:
 **************************************************************************/
func readIPFile() (src []string, err error) {
	f, e := os.Open("ip.txt")
	if e != nil {
		err = e
		return
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		src = append(src, line)
	}
	return
}

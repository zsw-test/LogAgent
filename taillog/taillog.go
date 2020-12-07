/*
 * @Author: your name
 * @Date: 2020-12-06 17:11:02
 * @LastEditTime: 2020-12-06 17:30:22
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \GoWork\LogAgent\taillog\taillog.go
 */
package tailLog

import (
	"fmt"

	"github.com/hpcloud/tail"
)

// 收集日志文件的模块

var (
	tailObj *tail.Tail
	LogChan chan string
)

func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed ,err :", err)
		return
	}
	return err
}

func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}

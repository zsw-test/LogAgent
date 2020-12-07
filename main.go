package main

import (
	"LogAgent/conf"
	"LogAgent/kafka"
	"LogAgent/tailLog"
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConfig)
)

func run() {
	// 读取日志 发送到kafka
	for {
		select {
		case line := <-tailLog.ReadChan():
			//发送到kafka
			kafka.SendToKafka(cfg.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

// logAgent 程序入口
func main() {
	//加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load config.ini failed, err:", err)
		return
	}
	//1：初始化kafka连接
	err = kafka.Init([]string{cfg.Address})
	if err != nil {
		fmt.Println("Init Kafka failed, err:", err)
		return
	}
	fmt.Println("init kafka success")
	//2：打开日志文件准备收集日志
	err = tailLog.Init(cfg.FileName)
	if err != nil {
		fmt.Println("Init Tail failed, err:", err)
		return
	}
	fmt.Println("init tailLog success")
	run()
}

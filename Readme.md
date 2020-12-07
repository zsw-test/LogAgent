首先配置kafka和zookeeper 的日志信息的保存位置

kafka依赖于zookeeper来管理集群borker各个结点的信息，服务注册发现等等

在windows下面要设置成正确的位置

```bash
kafka:
# A comma separated list of directories under which to store log files
log.dirs=E:/zsw/kafka-logs

zookeeper:
# the directory where the snapshot is stored.
dataDir=E:/zsw/zookeeper

```

然后通过管理员模式启动zookeeper

```bash
bin\windows\zookeeper-server-start.bat config\zookeeper.properties
```

再启动kafka

```bash
bin\windows\kafka-server-start.bat config\server.properties
```

然后就可以启动logagent.exe 来转发日志my.log信息发送到kafka集群

配置config.ini

```bash
[kafka]
address=127.0.0.1:9092
topic=web_log

[taillog]
path=./my.log
```

最后可以使用kafka自带的console-consumer来消费指定topic的信息

```bash
kafka-console-consumer.bat --bootstrap-server=127.0.0.1:9092 --from-beginning --topic web_log
```

然后可以读取消费kafka里面的指定topic的日志


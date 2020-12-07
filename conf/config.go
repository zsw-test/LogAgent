package conf

type AppConfig struct {
	KafkaConf   `ini:"kafka"`
	TaillogConf `ini:"taillog"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type TaillogConf struct {
	FileName string `ini:"path"`
}

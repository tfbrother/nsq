package nsqlookupd

import (
	"log"
	"os"
	"time"

	"github.com/tfbrother/nsq/internal/lg"
)

type Options struct {
	LogLevel  string `flag:"log-level"`
	LogPrefix string `flag:"log-prefix"`
	Verbose   bool   `flag:"verbose"` // for backwards compatibility 是否允许输出日志
	Logger    Logger
	logLevel  lg.LogLevel // private, not really an option

	TCPAddress  string `flag:"tcp-address"`
	HTTPAddress string `flag:"http-address"`
	// 这个lookupd节点的外部地址
	BroadcastAddress string `flag:"broadcast-address"`

	// 从上次 ping 之后，生产者驻留在活跃列表中的时长
	InactiveProducerTimeout time.Duration `flag:"inactive-producer-timeout"`
	// 生产者逻辑删除后，保持的时长
	TombstoneLifetime time.Duration `flag:"tombstone-lifetime"`
}

func NewOptions() *Options {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	return &Options{
		LogPrefix:        "[nsqlookupd] ",
		LogLevel:         "debug",
		TCPAddress:       "0.0.0.0:4160",
		HTTPAddress:      "0.0.0.0:4161",
		BroadcastAddress: hostname,

		InactiveProducerTimeout: 300 * time.Second,
		TombstoneLifetime:       45 * time.Second,
	}
}

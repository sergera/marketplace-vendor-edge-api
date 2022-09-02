package conf

import (
	"log"
	"sync"

	"github.com/gurkankaymak/hocon"
)

var once sync.Once
var instance *conf

type conf struct {
	hocon      *hocon.Config
	Port       string
	WorkerHost string
	WorkerPort string
}

func GetConf() *conf {
	once.Do(func() {
		var c *conf = &conf{}
		c.setup()
		instance = c
	})
	return instance
}

func (c *conf) setup() {
	c.parseHOCONConfigFile()
	c.setPort()
	c.setWorkerHost()
	c.setWorkerPort()
}

func (c *conf) parseHOCONConfigFile() {
	hocon, err := hocon.ParseResource("application.conf")
	if err != nil {
		log.Panic("error while parsing configuration file: ", err)
	}

	log.Printf("configurations: %+v", *hocon)

	c.hocon = hocon
}

func (c *conf) setPort() {
	port := c.hocon.GetString("host.port")
	if len(port) == 0 {
		log.Panic("port environment variable not found")
	}

	c.Port = port
}

func (c *conf) setWorkerHost() {
	workerHost := c.hocon.GetString("db.host")
	if len(workerHost) == 0 {
		log.Panic("worker host environment variable not found")
	}

	c.WorkerHost = workerHost
}

func (c *conf) setWorkerPort() {
	workerPort := c.hocon.GetString("db.port")
	if len(workerPort) == 0 {
		log.Panic("worker port environment variable not found")
	}

	c.WorkerPort = workerPort
}

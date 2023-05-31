package config

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type ConfigStruct struct {
	Debug       bool        `toml:"debug"`
	DB          DB          `toml:"db"`
	CoreService GRPCService `toml:"core"`
	NodeService GRPCService `toml:"node"`
	QuicService QuicService `toml:"quic"`
	Status      Status      `toml:"status"`
	Gin         Gin         `toml:"gin"`
	HttpService HttpService `toml:"http"`
	Statics     []Static    `toml:"static"`
}

type DB struct {
	Debug bool   `toml:"debug"`
	File  string `toml:"file"`
}

type GRPCService struct {
	Enable bool   `toml:"enable"`
	Addr   string `toml:"addr"`
	TLS    bool   `toml:"tls"`
	CA     string `toml:"ca"`
	Cert   string `toml:"cert"`
	Key    string `toml:"key"`
}

type Status struct {
	LinkTTL int `toml:"link_ttl"`
}

type QuicService struct {
	Enable bool   `toml:"enable"`
	Addr   string `toml:"addr"`
	CA     string `toml:"ca"`
	Cert   string `toml:"cert"`
	Key    string `toml:"key"`
}

type Gin struct {
	Debug bool `toml:"debug"`
}

type HttpService struct {
	Enable bool   `toml:"enable"`
	Addr   string `toml:"addr"`
	TLS    bool   `toml:"tls"`
	Cert   string `toml:"cert"`
	Key    string `toml:"key"`
}

type Static struct {
	Addr string `toml:"addr"`
	Path string `toml:"path"`
	TLS  bool   `toml:"tls"`
	Cert string `toml:"cert"`
	Key  string `toml:"key"`
}

func DefaultConfig() ConfigStruct {
	return ConfigStruct{
		Debug: false,
		DB: DB{
			File: "store.db",
		},
		CoreService: GRPCService{
			Enable: true,
			Addr:   ":6006",
			TLS:    true,
			CA:     "certs/ca.crt",
			Cert:   "certs/server.crt",
			Key:    "certs/server.key",
		},
		NodeService: GRPCService{
			Enable: true,
			Addr:   ":6007",
			TLS:    true,
			CA:     "certs/ca.crt",
			Cert:   "certs/server.crt",
			Key:    "certs/server.key",
		},
		QuicService: QuicService{
			Enable: true,
			Addr:   ":6008",
			CA:     "certs/ca.crt",
			Cert:   "certs/server.crt",
			Key:    "certs/server.key",
		},
		Status: Status{
			LinkTTL: 3 * 60,
		},
	}
}

func (c *ConfigStruct) Validate() error {
	if len(c.DB.File) == 0 {
		return errors.New("DB.File must be specified")
	}

	if c.CoreService.Enable {
		if len(c.CoreService.Addr) == 0 {
			return errors.New("CoreService.Addr must be specified")
		}

		if c.CoreService.TLS {
			if len(c.CoreService.CA) == 0 {
				return errors.New("CoreService.CA must be specified")
			}

			if len(c.CoreService.Cert) == 0 {
				return errors.New("CoreService.Cert must be specified")
			}

			if len(c.CoreService.Key) == 0 {
				return errors.New("CoreService.Key must be specified")
			}
		}
	}

	if c.NodeService.Enable {
		if len(c.NodeService.Addr) == 0 {
			return errors.New("NodeService.Addr must be specified")
		}

		if c.NodeService.TLS {
			if len(c.NodeService.CA) == 0 {
				return errors.New("NodeService.CA must be specified")
			}

			if len(c.NodeService.Cert) == 0 {
				return errors.New("NodeService.Cert must be specified")
			}

			if len(c.NodeService.Key) == 0 {
				return errors.New("NodeService.Key must be specified")
			}
		}
	}

	if c.QuicService.Enable {
		if len(c.QuicService.Addr) == 0 {
			return errors.New("QuicService.Addr must be specified")
		}

		if len(c.QuicService.CA) == 0 {
			return errors.New("QuicService.CA must be specified")
		}

		if len(c.QuicService.Cert) == 0 {
			return errors.New("QuicService.Cert must be specified")
		}

		if len(c.QuicService.Key) == 0 {
			return errors.New("QuicService.Key must be specified")
		}
	}

	return nil
}

var Config = DefaultConfig()

func Parse() {
	var err error

	configFile := flag.String("c", "config.toml", "config file")
	flag.Parse()

	if _, err = toml.DecodeFile(*configFile, &Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = Config.Validate(); err != nil {
		fmt.Println("config:", err)
		os.Exit(1)
	}
}
